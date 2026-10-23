package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int16, p2 int16) bool {
    if (!(!true)) {
        var v4 uint8 = 42
        _ = v4
    } else {
        var v5 int = -1
        p2 -= p1
        var v6 uint64 = 1
        _ = v5
        _ = v6
    }
    var v7 uint = uint(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (!true)
}
func genFn1(p0 uint16) uint64 {
    var v8 uint16 = p0
    if (true || false) {
        switch 4294967295 {
        case 32768:
            var v9 uint8 = 1
            _ = v9
        case 1:
            var v10 int64 = 9223372036854775807
            _ = v10
        case 128:
            v8 -= (v8 << ((966) & 63))
        }
    } else {
        p0 = v8
    }
    var v11 int64 = 1
    _ = p0
    _ = v11
    _ = v8
    return 244
}

func main() {
    var v1 int = 1
    v2 := []int{128, v1, (v1 ^ (v1 & 9223372036854775807))}
    var v3 float32 = (((1.5 / 0.1) + (1.5 - (3.14159 / 1e10))) - (0.0 * (float32(1.5) * float32(3.14159))))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v12 uint64 = 2
    if (-976 >= 2147483647) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        for fi13 := 0; fi13 < 3; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(fi14, v12)
                _ = fi14
            }
            v1 = (fi13 - (v1 & fi13))
            _ = fi13
        }
    }
    var v15 int32 = 42
    switch (-(v1)) {
    case -123:
        v2 = append(v2, v1)
    case 0:
        var v16 uint8 = 255
        _ = v16
    default:
        switch ((("foo bar" + "") + "hello") + ("\n\t" + "")) {
        case "\n\t":
            v3 += v3
        case "a":
            var v17 bool = ("hello" < (((("foo bar" + "a") + (("hello" + "") + ("\n\t" + "\n\t"))) + "") + "hello"))
            _ = v17
        default:
            v12 -= (v12 << ((1) & 63))
        }
    }
    switch ("hello" + "") {
    case "\n\t":
        v15 -= (v15 * (v15 ^ (v15 * (v15 >> ((326) & 63)))))
    case "":
        for fi18 := 0; fi18 < 14; fi18++ {
            switch (("a" + "foo bar") + ("" + ("" + "hello"))) {
            case "hello":
                var v19 int = (v1 & (^(fi18)))
                _ = v19
            default:
                v3 -= v3
            }
            v12 -= genFn1(32767)
            fmt.Println(v1)
            _ = fi18
        }
    default:
        v20 := GenS1{F0: 42}
        _ = v20
    }
    var v21 int32 = v15
    v22 := GenS1{F0: 32767}
    v23 := GenS1{F0: 1}
    fmt.Println(v21)
    var v24 int8 = -1
    v25 := GenS1{F0: 255}
    fmt.Println(v15)
    fmt.Println(v12, v1)
    var v26 bool = ((false && true) && (-9223372036854775808 > 0))
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 int32 = 1
        v3 = (((v3 * v3) + ((1.5 / v3) / 0.1)) + v3)
        _ = fi27
        _ = v28
    }
    v2 = append(v2, 495)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
