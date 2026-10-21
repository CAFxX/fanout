package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int16; F2 int }
func genFn0(p0 string, p1 int16, p2 uint) int8 {
    v4 := []int{-1, -9223372036854775808}
    var v5 float32 = (-2.25 * (1.5 * -2.25))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"\n\t"}
    var v3 float32 = 0.1
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(fi6, v3)
        v3 = 1.5
        _ = fi6
    }
    v3 -= v3
    fmt.Println(v3)
    if true {
        _ = len(v2)
        if false {
            for fi7 := 0; fi7 < 2; fi7++ {
                v3 = float32(1.5)
                v3 += float32(float64(float32(1.5)))
                _ = fi7
            }
            var v8 string = "hello"
            _ = len(v1)
            _ = v8
        }
    } else {
        v3 = ((v3 - v3) + v3)
    }
    v3 += v3
    switch 2147483647 {
    case 0:
        v3 -= (v3 - 0.0)
    case -2147483648:
        v9 := GenS1{F0: 552}
        _ = v9
    case 127:
        if true {
            v3 += v3
            var v10 uint = 646
            v11 := v1[2147483648]
            _ = v10
            _ = v11
        }
    }
    switch ((("a" + ("" + ("" + "a"))) + "a") + "a") {
    case "hello":
        var v12 string = "a"
        _ = v12
    case "foo bar":
        if true {
            v13 := GenS2{F0: 2147483647, F1: 0, F2: 0}
            _ = v13
        } else {
            var v14 float64 = ((-2.25 * (-2.25 / 1.0)) - (0.0 - 1e10))
            v3 -= (v3 + (0.0 * (0.1 - v3)))
            _ = v14
        }
    default:
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    if (!true) {
        var v15 uint8 = 2
        _ = v15
    }
    if (!(1 != 0)) {
        var v16 int8 = 100
        var v17 int = 32767
        if (!true) {
            v16 += genFn0("a", 32767, 926)
        } else {
            for fi18 := 0; fi18 < 6; fi18++ {
                for fi19 := 0; fi19 < 16; fi19++ {
                    switch (v17 * (fi18 & (v17 * 2147483647))) {
                    case 2147483647:
                        v1[1] = "\n\t"
                    default:
                        v16 = v16
                    }
                    _ = fi19
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v20 := GenS2{F0: 2147483647, F1: 1000, F2: (fi18 >> ((4294967295) & 63))}
                _ = fi18
                _ = v20
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 bool = ((423 >= 980) && (((fi21 + (v17 - -2147483648)) != (v17 % ((fi21) | 1))) && true))
                if len(v2) > 0 {
                    v2[0] = ("\n\t" + "hello")
                }
                _ = fi21
                _ = v22
            }
            v23 := GenS2{F0: 1, F1: -1, F2: -1}
            _ = v23
        }
        _ = v16
        _ = v17
    } else {
        var v24 uint8 = 0
        fmt.Println(v3, v24)
        if true {
            switch "" {
            case "a":
                switch -478 {
                case 0:
                    var v25 bool = false
                    _ = v25
                default:
                    var v26 int = 2147483647
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 19; fi27++ {
                switch ("a" + ("hello" + "")) {
                case "":
                    var v28 string = "\n\t"
                    _ = v28
                case "hello":
                    switch (fi27 % ((-34) | 1)) {
                    case 127:
                        var v29 uint32 = 42
                        _ = v29
                    case 32768:
                        for fi30 := 0; fi30 < 4; fi30++ {
                            switch (fi30 / ((fi27) | 1)) {
                            case 42:
                                v2 = append(v2, "")
                            }
                            _ = fi30
                        }
                    default:
                        var v31 float32 = v3
                        _ = v31
                    }
                default:
                    var v32 float32 = 3.14159
                    _ = v32
                }
                _ = fi27
            }
            for fi33 := 0; fi33 < 20; fi33++ {
                v24 += v24
                v34 := GenS2{F0: 42, F1: 1, F2: (fi33 - (fi33 & fi33))}
                _ = fi33
                _ = v34
            }
        }
        _ = v24
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        switch fi35 {
        case 2147483647:
            v3 -= (v3 - (v3 + (v3 - -2.25)))
        case 32768:
            v2 = append(v2, ("hello" + ("a" + ("\n\t" + (((("\n\t" + "foo bar") + ("foo bar" + ("hello" + "a"))) + "foo bar") + "")))))
        case 4294967295:
            v1[1] = ""
        }
        _ = fi35
    }
    if (!true) {
        v3 += 1e10
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
