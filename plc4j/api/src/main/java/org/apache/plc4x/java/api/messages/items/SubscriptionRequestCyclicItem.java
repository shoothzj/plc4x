/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.api.messages.items;

import org.apache.plc4x.java.api.model.Address;
import org.apache.plc4x.java.api.model.SubscriptionType;

import java.util.Objects;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

public class SubscriptionRequestCyclicItem<T> extends SubscriptionRequestItem<T> {

    private TimeUnit timeUnit;
    private int period;

    public SubscriptionRequestCyclicItem(Class<T> dataType, Address address, TimeUnit timeUnit, int period, Consumer<SubscriptionEventItem<T>> consumer) {
        super(dataType, address, SubscriptionType.CYCLIC, consumer);
        this.timeUnit = timeUnit;
        this.period = period;
    }

    public TimeUnit getTimeUnit() {
        return timeUnit;
    }

    public int getPeriod() {
        return period;
    }

    @Override
    public String toString() {
        return "SubscriptionRequestCyclicItem{" +
            "timeUnit=" + timeUnit +
            ", period=" + period +
            "} " + super.toString();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (!(o instanceof SubscriptionRequestCyclicItem)) {
            return false;
        }
        if (!super.equals(o)) {
            return false;
        }
        SubscriptionRequestCyclicItem that = (SubscriptionRequestCyclicItem) o;
        return period == that.period &&
            timeUnit == that.timeUnit;
    }

    @Override
    public int hashCode() {
        return Objects.hash(super.hashCode(), timeUnit, period);
    }
}
