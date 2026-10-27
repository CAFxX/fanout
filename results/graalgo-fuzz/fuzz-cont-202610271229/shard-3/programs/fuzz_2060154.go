package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8) string {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 float64 = 1.5
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint = 1
        _ = fi7
        _ = v8
    }
    _ = p0
    return ("\n\t" + ("" + "hello"))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 210
    var v3 int = -944
    v4 := make(map[uint32]bool)
    var v9 int32 = 0
    v2 = (^(v2))
    fmt.Println(v9, v9)
    v2 += v2
    var v10 uint64 = 303
    fmt.Println(v9)
    var v11 float64 = (((((-2.25 * 3.14159) - (-2.25 - float64(3.14159))) * 0.0) - 0.1) * (0.1 * -2.25))
    var v12 uint16 = 38
    switch (("\n\t" + ("foo bar" + ((((("\n\t" + "foo bar") + ("hello" + "hello")) + ("\n\t" + ("hello" + ("" + ((("a" + "") + "hello") + "a"))))) + "a") + ""))) + "") {
    case "":
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint8 = v2
            var v15 int16 = int16(v14)
            if (!true) {
                for fi16 := 0; fi16 < 10; fi16++ {
                    v10 = v10
                    _ = fi16
                }
                fmt.Println(fi13, v9)
                v17 := GenS1{F0: 42}
                _ = v17
            } else {
                fmt.Println(v11, v14, v11)
                delete(v1, fi13)
                for fi18 := 0; fi18 < 13; fi18++ {
                    v19 := GenS1{F0: 42}
                    _ = fi18
                    _ = v19
                }
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    case "hello":
        switch "\n\t" {
        case "":
            var v20 uint8 = (v2 * (v2 * (^(v2))))
            _ = v20
        case "foo bar":
            if true {
                v21 := []int{588, v3, 2147483647}
                delete(v4, uint32(v12))
                _ = len(v4)
                _ = v21
            }
        default:
            v11 += (0.0 / 1.5)
        }
    }
    v10 = (v10 / ((v10) | 1))
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            v11 += v11
            v24 := GenS1{F0: 127}
            _ = fi23
            _ = v24
        }
        if true {
            v25 := GenS2{F0: 91}
            v9 = 2147483647
            _ = v25
        } else {
            switch "hello" {
            case "hello":
                if (-128 >= -1) {
                    v2 = 184
                    var v26 int8 = 127
                    fmt.Println(v3, v10)
                    _ = v26
                } else {
                    v27 := GenS1{F0: -41}
                    _ = v27
                }
            case "":
                for fi28 := 0; fi28 < 3; fi28++ {
                    switch "\n\t" {
                    case "":
                        var v29 uint32 = 410
                        _ = v29
                    case "foo bar":
                        v11 = 1e10
                    }
                    var v30 uint64 = (v10 / ((v10) | 1))
                    fmt.Println(v30)
                    _ = fi28
                    _ = v30
                }
            case "\n\t":
                if (false && (729 == 0)) {
                    if (!((("a" + "hello") != "") || false)) {
                        v31 := GenS2{F0: 42}
                        if (!(v9 >= v9)) {
                            v32 := []int{127, v3, (fi22 / ((fi22) | 1))}
                            if len(v32) > 0 {
                                v32[0] = (^(fi22))
                            }
                            _ = v32
                        } else {
                            v33 := GenS1{F0: -48}
                            for fi34 := 0; fi34 < 17; fi34++ {
                                v10 -= (v10 + v10)
                                _ = fi34
                            }
                            _ = v33
                        }
                        _ = v31
                    }
                    if true {
                        var v35 int16 = 1
                        _ = v35
                    } else {
                        if (v12 == v12) {
                            var v36 uint64 = (v10 | uint64(v2))
                            var v37 uint64 = v36
                            v37 += v37
                            _ = v36
                            _ = v37
                        } else {
                            v38 := GenS1{F0: -128}
                            _ = v38
                        }
                    }
                }
            }
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
