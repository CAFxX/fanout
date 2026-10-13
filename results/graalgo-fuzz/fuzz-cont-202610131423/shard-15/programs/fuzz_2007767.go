package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 1
    v3 := []uint32{591, 980, 966, 375}
    var v4 uint32 = 715
    v5 := []int{0}
    v6 := GenS1{F0: -9223372036854775808}
    var v7 bool = (-1 > 459)
    fmt.Println(v7, v7, v2)
    switch 4294967295 {
    case -1:
        for fi8 := 0; fi8 < 2; fi8++ {
            switch -1 {
            case -842:
                var v9 int32 = (^(v2))
                _ = v9
            case 83:
                for fi10 := 0; fi10 < 9; fi10++ {
                    _ = len(v1)
                    if len(v3) > 0 {
                        v3[0] = v4
                    }
                    if (255 <= 65535) {
                        v11 := GenS1{F0: -9223372036854775808}
                        var v12 int8 = 100
                        _ = v11
                        _ = v12
                    }
                    _ = fi10
                }
            }
            v13 := GenS1{F0: -1}
            v14 := GenS1{F0: -1}
            _ = fi8
            _ = v13
            _ = v14
        }
    case 32768:
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS1{F0: -513}
            if (2 <= 1) {
                if v7 {
                    v1[("hello" + "foo bar")] = ""
                }
            } else {
                for fi17 := 0; fi17 < 5; fi17++ {
                    v5 = append(v5, fi17)
                    fmt.Println(v7)
                    var v18 float64 = 3.14159
                    if (((fi17 - fi15) < fi17) && (v2 < -2147483648)) {
                        fmt.Println(v4)
                    } else {
                        v19 := v1["\n\t"]
                        _ = v19
                    }
                    v20 := GenS1{F0: -9223372036854775808}
                    _ = fi17
                    _ = v18
                    _ = v20
                }
            }
            _ = fi15
            _ = v16
        }
    default:
        v21 := GenS2{F0: 42}
        _ = v21
    }
    v2 = v2
    if (!((735 == -32768) && (!false))) {
        v1[("foo bar" + "\n\t")] = "a"
        v7 = (1 != 288)
        v2 = (-(v2))
    }
    fmt.Println(v7, v7)
    switch int(v2) {
    case 1:
        v1["foo bar"] = (("" + ("" + "foo bar")) + ("foo bar" + ""))
    }
    fmt.Println(v7)
    v4 -= (v4 & v4)
    if (!v7) {
        v22 := GenS1{F0: -1}
        v4 += v4
        if len(v5) > 0 {
            v5[0] = -959
        }
        _ = v22
    } else {
        if v7 {
            v23 := GenS1{F0: 0}
            var v24 float32 = 1e10
            if false {
                var v25 uint16 = 42
                v2 += v2
                _ = v25
            } else {
                v2 = v2
                var v26 uint8 = 127
                _ = v26
            }
            _ = v23
            _ = v24
        } else {
            v7 = v7
            var v27 uint32 = 4294967295
            _ = v27
        }
    }
    if (!(((4294967295 != 1) || (0.0 == (((((((-2.25 - 1.5) + 1.5) * ((0.1 - float64(1.5)) - (((1.5 / -2.25) / (1.5 - 0.1)) + ((-2.25 / (((3.14159 / -2.25) - 1.5) / 3.14159)) - -2.25)))) / (-2.25 - 1e10)) - 0.0) - -2.25) + -2.25))) && ("hello" != "foo bar"))) {
        v7 = false
    }
    switch -1 {
    case 276:
        v7 = false
    case 0:
        v28 := GenS1{F0: 0}
        _ = v28
    case 32767:
        v5 = append(v5, 32767)
    }
    v29 := GenS2{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
