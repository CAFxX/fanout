package main

import (
    "fmt"
)

func genFn0() uint8 {
    v4 := []int{32768}
    v4 = append(v4, 9223372036854775807)
    switch -9223372036854775808 {
    case 1:
        if (-420 == 4294967295) {
            var v5 uint64 = 4294967295
            v5 -= v5
            _ = v5
        } else {
            var v6 int = -138
            var v7 uint16 = 895
            if len(v4) > 0 {
                v4[0] = v6
            }
            _ = v6
            _ = v7
        }
    }
    var v8 int16 = 1000
    var v9 uint32 = 0
    _ = v4
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int8) int32 {
    var v10 uint64 = 42
    var v11 uint8 = 0
    _ = p0
    _ = v10
    _ = v11
    return int32(v11)
}

func main() {
    var v1 bool = false
    v2 := []string{"hello", "hello"}
    v3 := []int{1, 4294967295}
    switch ("foo bar" + (("foo bar" + "a") + "foo bar")) {
    case "":
        v1 = v1
    case "\n\t":
        _ = len(v2)
    case "foo bar":
        v1 = false
    }
    var v12 uint64 = 42
    v1 = v1
    v1 = v1
    if (193 == 42) {
        v1 = ((-171 <= 42) && v1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v13 int = 653
    for fi14 := 0; fi14 < 17; fi14++ {
        switch (-(v13)) {
        case 32768:
            v2 = append(v2, ("" + "a"))
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint64 = v12
            _ = fi15
            _ = v16
        }
        if v1 {
            if (((-1 != 1) || (!(-32768 < -1))) && (!false)) {
                v1 = true
                if ((542 <= 0) && (float32(-2.25) <= 0.1)) {
                    var v17 int8 = 100
                    v17 = (-(v17))
                    v13 = 871
                    _ = v17
                }
            }
        } else {
            v1 = v1
            var v18 uint32 = 65535
            _ = v18
        }
        _ = fi14
    }
    v13 = v13
    if v1 {
        fmt.Println(v1)
    } else {
        v13 -= (-(v13))
        var v19 uint64 = (^(v12))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
