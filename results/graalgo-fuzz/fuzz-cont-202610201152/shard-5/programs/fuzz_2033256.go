package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint64, p1 int16) bool {
    v6 := GenS1{F0: -712}
    if false {
        var v7 uint64 = 0
        switch 32768 {
        case 4294967295:
            var v8 int8 = 1
            _ = v8
        case 829:
            var v9 int64 = -1
            _ = v9
        case -611:
            var v10 uint8 = 2
            _ = v10
        default:
            var v11 uint32 = 42
            _ = v11
        }
        p1 = p1
        _ = v7
    }
    if (!false) {
        var v12 int8 = 42
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    return true
}

func main() {
    v1 := []uint32{42, 2147483647}
    var v2 int16 = -1
    var v3 int = 4294967295
    v4 := []string{""}
    var v5 bool = (168 != 2)
    for fi13 := 0; fi13 < 2; fi13++ {
        switch "a" {
        case "foo bar":
            for fi14 := 0; fi14 < 8; fi14++ {
                switch (v3 + (^(fi14))) {
                case -352:
                    v3 = fi14
                case 4294967295:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 2147483647:
                    fmt.Println(v2)
                    if (20 >= -1) {
                        _ = len(v1)
                        v2 -= (v2 % ((v2) | 1))
                        fmt.Println(fi14, fi13, v3)
                        _ = len(v1)
                    } else {
                        v4 = append(v4, ("" + ("a" + ("\n\t" + ""))))
                        if (true && v5) {
                            v5 = (1e10 > (((0.0 * ((0.1 + (1.5 * (-2.25 / float32(1.5)))) * (1.5 - 1.5))) + (3.14159 / 0.1)) - 0.0))
                            v3 -= (v3 * (fi13 ^ fi14))
                        }
                        var v15 int64 = 42
                        _ = v15
                    }
                }
                _ = fi14
            }
        case "":
            var v16 int32 = -209
            _ = v16
        case "\n\t":
            if (true && v5) {
                for fi17 := 0; fi17 < 3; fi17++ {
                    _ = len(v4)
                    fmt.Println(fi17, v3, fi17)
                    if v5 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = fi17
                }
            }
        }
        v18 := GenS1{F0: 1}
        if false {
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 string = ("" + "hello")
                var v21 float64 = -2.25
                if v5 {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    v3 += (v3 / (((^(fi19))) | 1))
                } else {
                    if (-878 >= -1) {
                        switch fi19 {
                        case -1:
                            var v22 int = 127
                            _ = v22
                        }
                    } else {
                        var v23 float32 = 0.0
                        fmt.Println(v23, v21)
                        _ = v23
                    }
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        _ = fi13
        _ = v18
    }
    v5 = v5
    for fi24 := 0; fi24 < 9; fi24++ {
        if true {
            v25 := GenS2{F0: 42}
            _ = v25
        }
        var v26 bool = v5
        _ = fi24
        _ = v26
    }
    var v27 int64 = 42
    var v28 uint16 = 566
    v29 := GenS2{F0: -128}
    if false {
        v28 = v28
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 964
        }
        if (v5 || ((!(!(1 < 0))) || v5)) {
            _ = len(v4)
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 float32 = ((float32((1e10 + 0.1)) + ((-2.25 + -2.25) * 1e10)) + 0.1)
                switch int(v27) {
                case 1:
                    _ = len(v4)
                default:
                    var v32 string = "a"
                    _ = v32
                }
                var v33 int64 = -9223372036854775808
                _ = fi30
                _ = v31
                _ = v33
            }
        }
    }
    var v34 uint8 = 2
    var v35 int8 = -1
    v36 := GenS1{F0: 0}
    v37 := GenS1{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
