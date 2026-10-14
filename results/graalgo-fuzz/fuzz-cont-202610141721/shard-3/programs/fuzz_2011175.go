package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
func genFn0() int16 {
    var v7 int16 = 994
    v7 = (v7 | (^(v7)))
    var v8 int8 = -71
    v9 := []int{-1, 2147483647}
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    var v1 int32 = -1
    v2 := []string{"hello", "a", "hello"}
    var v3 int16 = 1000
    var v4 string = "hello"
    v5 := make(map[int]uint32)
    v6 := make(map[string]bool)
    if true {
        v10 := v6[("hello" + v4)]
        switch (((((("hello" + v4) + ("foo bar" + v4)) + "foo bar") + (v4 + v4)) + (v4 + (v4 + v4))) + ((v4 + "hello") + v4)) {
        case "a":
            v3 -= genFn0()
        case "":
            fmt.Println(v4, v4)
        case "hello":
            var v11 string = ((v4 + v4) + "foo bar")
            _ = v11
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            v3 += (v3 << ((255) & 63))
            var v13 float64 = -2.25
            v3 += genFn0()
            _ = fi12
            _ = v13
        }
        _ = v10
    } else {
        switch 42 {
        case 32767:
            fmt.Println(v4, v4, v3)
        }
        fmt.Println(v3, v3, v4)
    }
    v14 := v6["\n\t"]
    if v14 {
        v15 := GenS1{F0: -1}
        v14 = v14
        _ = v15
    } else {
        switch v4 {
        case "foo bar":
            var v16 bool = (((1.5 * 0.0) > (0.0 / 0.1)) && ((((v4 + "hello") + v4) < v4) || (false || v14)))
            _ = v16
        case "a":
            v5[-9223372036854775808] = 65535
        case "":
            switch -586 {
            case 128:
                v17 := GenS1{F0: -1}
                _ = v17
            case -291:
                v4 = v4
            case 127:
                if v14 {
                    fmt.Println(v1, v14)
                    var v18 uint8 = 44
                    if (v4 <= (((v4 + v4) + v4) + (v4 + (v4 + "foo bar")))) {
                        v1 += -1
                        switch v4 {
                        case "hello":
                            v19 := GenS1{F0: -86}
                            _ = v19
                        }
                        var v20 bool = v14
                        _ = v20
                    }
                    _ = v18
                }
            }
        default:
            var v21 bool = ((v14 && v14) || (42 == 1))
            _ = v21
        }
    }
    v5[127] = 1
    if v14 {
        if false {
            var v22 int64 = -1
            v4 = (v4 + v4)
            _ = v22
        }
    }
    fmt.Println(v1, v14)
    v4 = (((("" + ("foo bar" + v4)) + v4) + "") + ("" + v4))
    v3 -= (v3 % ((-32768) | 1))
    v14 = (!false)
    v3 -= v3
    for fi23 := 0; fi23 < 11; fi23++ {
        fmt.Println(v1, v3, v4)
        _ = fi23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 4; fi25++ {
            if v14 {
                if ((!(v4 != v4)) && v14) {
                    v3 = (v3 << ((374) & 63))
                    fmt.Println(fi25, v14)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v26 float32 = float32(1e10)
                _ = v26
            } else {
                v27 := GenS2{F0: 100}
                _ = v27
            }
            v28 := GenS2{F0: -27}
            switch fi25 {
            case 128:
                switch (((v4 + v4) + v4) + ("" + v4)) {
                case "hello":
                    if (3.14159 == (0.1 * (float64(0.1) + 1e10))) {
                        v3 = -32768
                    }
                case "foo bar":
                    fmt.Println(v1, v3)
                case "":
                    v29 := GenS1{F0: 73}
                    _ = v29
                }
            case -552:
                var v30 int64 = 9223372036854775807
                _ = v30
            default:
                _ = len(v5)
            }
            _ = fi25
            _ = v28
        }
        _ = fi24
    }
    v31 := GenS2{F0: 88}
    v14 = (2 < 127)
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
