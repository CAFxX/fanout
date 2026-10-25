package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int8 {
    var v7 int = -2147483648
    var v8 int64 = 1
    v9 := GenS1{F0: 423}
    if (!((((3.14159 - (0.0 - (float64((float32(0.1) / 1.5)) - (1.5 * 3.14159)))) * -2.25) / 1e10) > 1e10)) {
        v7 -= v7
    } else {
        switch (^(v7)) {
        case 1:
            var v10 uint = 4294967295
            _ = v10
        }
        v11 := []int{89}
        _ = v11
    }
    var v12 bool = (v7 == (v7 | 77))
    _ = v12
    _ = v7
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 uint8 = 42
    var v2 uint = 42
    var v3 uint32 = 1
    v4 := []int{127}
    var v5 uint16 = 255
    var v6 uint16 = v5
    for fi13 := 0; fi13 < 20; fi13++ {
        switch "foo bar" {
        case "foo bar":
            switch (((("\n\t" + "hello") + "hello") + "a") + (("foo bar" + "a") + (("foo bar" + "\n\t") + "\n\t"))) {
            case "a":
                for fi14 := 0; fi14 < 3; fi14++ {
                    v1 = (^(v1))
                    _ = fi14
                }
            case "\n\t":
                v3 = v3
            case "hello":
                if len(v4) > 0 {
                    v4[0] = fi13
                }
            }
        case "":
            v5 += v5
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = len(v4)
        _ = len(v4)
        _ = fi13
    }
    v15 := GenS1{F0: 42}
    var v16 int16 = 714
    v6 -= v6
    var v17 uint8 = v1
    if true {
        switch 460 {
        case 102:
            if (3.14159 == (-2.25 - -2.25)) {
                for fi18 := 0; fi18 < 7; fi18++ {
                    if (1 > 2147483647) {
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v1 -= 42
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi19
                        }
                    }
                    v4 = append(v4, 9223372036854775807)
                    _ = fi18
                }
            }
        }
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        fmt.Println(v5, v3)
        if (true && (true && false)) {
            v2 += v2
        }
        if (3.14159 >= 1.5) {
            switch "foo bar" {
            case "hello":
                if (!(404 <= 2147483648)) {
                    var v21 int = fi20
                    if (-75 <= 0) {
                        var v22 string = "a"
                        _ = v22
                    }
                    if (false || (0 <= (v5 / ((65535) | 1)))) {
                        v6 = (v6 - v5)
                    }
                    _ = v21
                }
            case "":
                var v23 int32 = 0
                _ = v23
            default:
                var v24 uint = 833
                _ = v24
            }
            var v25 int8 = -60
            _ = v25
        } else {
            switch "a" {
            case "hello":
                if false {
                    fmt.Println(v1, v6)
                    var v26 int = (fi20 / ((fi20) | 1))
                    if len(v4) > 0 {
                        v4[0] = (v26 * 127)
                    }
                    _ = v26
                }
            default:
                var v27 int = (fi20 ^ (fi20 >> ((v2) & 63)))
                _ = v27
            }
        }
        _ = fi20
    }
    switch "hello" {
    case "hello":
        fmt.Println(v3, v16, v16)
    default:
        v6 = v6
    }
    var v28 string = ("\n\t" + "hello")
    var v29 int32 = int32(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
