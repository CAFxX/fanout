package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 bool }
func genFn0(p0 int32, p1 uint, p2 uint32) bool {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 15; fi5++ {
            p1 = p1
            p2 = (p2 * p2)
            _ = fi5
        }
        _ = fi4
    }
    var v6 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}

func main() {
    var v1 int16 = 1
    v2 := make(map[string]string)
    v3 := []int{9223372036854775807, -2147483648}
    var v7 uint = 42
    var v8 uint = 1
    delete(v2, "")
    v3 = append(v3, 74)
    fmt.Println(v8, v1, v8)
    var v9 float32 = float32(-2.25)
    for fi10 := 0; fi10 < 13; fi10++ {
        switch fi10 {
        case -382:
            switch "foo bar" {
            case "foo bar":
                if true {
                    switch "" {
                    case "hello":
                        v7 = (v7 >> ((v7) & 63))
                    default:
                        if false {
                            switch fi10 {
                            case 128:
                                var v11 uint16 = 2
                                _ = v11
                            default:
                                var v12 uint32 = 2147483647
                                _ = v12
                            }
                        } else {
                            _ = len(v3)
                            var v13 uint32 = 65535
                            _ = v13
                        }
                    }
                    v7 = (v7 | 936)
                }
            default:
                for fi14 := 0; fi14 < 19; fi14++ {
                    _ = len(v3)
                    v15 := v2[("\n\t" + "foo bar")]
                    _ = fi14
                    _ = v15
                }
            }
        case 2147483647:
            for fi16 := 0; fi16 < 17; fi16++ {
                v8 = 4294967295
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    delete(v2, ("hello" + "hello"))
                }
                var v17 int = 9223372036854775807
                var v18 int = (fi10 + 4294967295)
                var v19 uint8 = 127
                v7 += (^(v8))
                v3 = append(v3, v17)
                switch ("a" + "foo bar") {
                case "hello":
                    switch (v17 << ((v8) & 63)) {
                    case 0:
                        if true {
                            _ = len(v3)
                            v20 := GenS1{F0: 42, F1: v19, F2: (false || (v19 > v19))}
                            if ((v19 << (((v8 - v8)) & 63)) < v19) {
                                var v21 uint16 = 1
                                var v22 float64 = ((1.5 / (1e10 * 1.5)) + 0.1)
                                v1 -= 32767
                                _ = v21
                                _ = v22
                            }
                            _ = v20
                        }
                    case -986:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -979:
                        switch "\n\t" {
                        case "hello":
                            v9 = float32(((((0.0 * (1.5 - ((1.5 - ((0.1 + (((3.14159 - float64((((float32(0.1) * (v9 / ((float32((3.14159 / (-2.25 / (3.14159 + -2.25)))) + ((((float32(-2.25) + v9) * v9) - v9) - v9)) + v9))) * v9) - float32(1.5)))) + -2.25) - 1e10)) + 3.14159)) * -2.25))) - ((0.1 - float64(-2.25)) + 1e10)) / 0.1) * (1.5 - (0.0 - ((float64(0.1) + float64((3.14159 / 1e10))) / float64(((1e10 - v9) + float32((float64(float32(0.0)) / ((((0.1 - float64(float32((0.0 + -2.25)))) / float64(v9)) * -2.25) + 0.1))))))))))
                            fmt.Println(v18, fi10, v1)
                            if len(v3) > 0 {
                                v3[0] = (fi10 - (^(v17)))
                            }
                            fmt.Println(v8, v1)
                        case "foo bar":
                            var v23 int64 = 42
                            _ = v23
                        case "a":
                            v3 = append(v3, fi16)
                        default:
                            v17 = fi16
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch (-(fi10)) {
                            case 4294967295:
                                var v24 int = fi16
                                _ = v24
                            }
                        }
                    }
                }
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        case -1:
            var v25 uint8 = 255
            _ = v25
        }
        fmt.Println(v9, fi10)
        fmt.Println(v1)
        _ = fi10
    }
    v2["hello"] = "hello"
    v26 := GenS1{F0: 2147483647, F1: 142, F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
