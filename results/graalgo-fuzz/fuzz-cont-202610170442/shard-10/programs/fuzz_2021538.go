package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 int32, p1 uint16) uint64 {
    switch -288 {
    case -1:
        p1 -= 255
    case 128:
        for fi5 := 0; fi5 < 8; fi5++ {
            p1 = p1
            _ = fi5
        }
    case -9223372036854775808:
        p1 = p1
    }
    v6 := GenS1{F0: 65535, F1: false}
    var v7 float64 = ((((0.1 + (0.1 - float64((float32(3.14159) * 3.14159)))) / 0.1) * 1e10) * 1.5)
    var v8 string = "a"
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 uint32) string {
    v9 := GenS2{F0: 1, F1: 1}
    var v10 uint64 = 154
    switch 127 {
    case 127:
        if false {
            v11 := GenS2{F0: 315, F1: 127}
            v12 := []int{32767}
            _ = v11
            _ = v12
        }
    }
    var v13 uint32 = p0
    _ = p0
    _ = v10
    _ = v13
    _ = v9
    return "foo bar"
}

func main() {
    var v1 uint16 = 65535
    var v2 int8 = 127
    var v3 float32 = 1e10
    var v4 int32 = 1
    fmt.Println(v4, v1, v4)
    if false {
        v14 := GenS1{F0: 2147483648, F1: ("a" != "a")}
        fmt.Println(v4, v3, v4)
        if (!(0.0 >= (float64(1e10) * 0.1))) {
            v15 := []int{792, 9223372036854775807, 42}
            v4 = v4
            _ = v15
        }
        _ = v14
    } else {
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                v2 -= v2
                _ = fi17
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 uint16 = (v1 - v1)
                if (2147483647 != (fi16 * (^(fi16)))) {
                    if true {
                        v3 += v3
                        var v20 uint = 0
                        _ = v20
                    } else {
                        var v21 int = fi18
                        _ = v21
                    }
                }
                _ = fi18
                _ = v19
            }
            v22 := GenS2{F0: 255, F1: v2}
            _ = fi16
            _ = v22
        }
        if false {
            v3 -= v3
            if (v3 != (3.14159 + (v3 / 1e10))) {
                var v23 uint8 = 42
                _ = v23
            }
        }
        v24 := []int{127, 32768, 127, -910}
        _ = v24
    }
    var v25 uint64 = 4294967295
    var v26 int16 = 622
    for fi27 := 0; fi27 < 7; fi27++ {
        v28 := GenS2{F0: 42, F1: v2}
        v2 += -82
        v29 := []int{}
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int32 = 2147483647
    var v31 int16 = (^(v26))
    var v32 uint8 = 128
    v3 -= ((v3 * 0.0) - ((float32(0.1) / 1.5) * ((-2.25 - v3) + 0.1)))
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 string = "hello"
        if (v2 >= 127) {
            var v35 int8 = 50
            for fi36 := 0; fi36 < 5; fi36++ {
                switch -9223372036854775808 {
                case -2147483648:
                    fmt.Println(v32)
                case -22:
                    v37 := []int{(fi33 - fi33), 1}
                    _ = v37
                }
                v2 += (v35 >> ((18446744073709551615) & 63))
                var v38 uint32 = 2147483647
                _ = fi36
                _ = v38
            }
            _ = v35
        } else {
            v2 = v2
            var v39 int = fi33
            v40 := []int{128}
            _ = v39
            _ = v40
        }
        switch 32767 {
        case 4294967295:
            var v41 float64 = 3.14159
            _ = v41
        }
        _ = fi33
        _ = v34
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        v3 = 1e10
        switch (fi42 % (((fi42 * (fi42 + 32768))) | 1)) {
        case 1:
            switch -998 {
            case -9223372036854775808:
                v25 += v25
            }
        default:
            fmt.Println(v30, v30, v31)
        }
        _ = fi42
    }
    var v43 uint32 = 4294967295
    var v44 int8 = 0
    switch "a" {
    case "a":
        if ("\n\t" == "") {
            v45 := []int{}
            v26 = -32768
            _ = v45
        }
    default:
        v46 := []int{-9223372036854775808, -790, -1, 690}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v43
    _ = v44
}
