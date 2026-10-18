package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []string{"foo bar", "a", ""}
    v2 := make(map[string]int)
    var v3 int32 = 2147483647
    var v4 bool = false
    v5 := make(map[string]int)
    switch 922 {
    case 32768:
        if false {
            _ = len(v1)
        }
    case 127:
        v5["a"] = 42
    default:
        v6 := GenS1{F0: 720}
        _ = v6
    }
    v3 = (v3 | -1)
    v3 -= (-(v3))
    if v4 {
        fmt.Println(v4, v4)
        var v7 uint64 = 0
        if (v4 && (!(v4 && (-1 <= 1)))) {
            var v8 int8 = -1
            _ = v8
        }
        _ = v7
    }
    var v9 float32 = 0.1
    v4 = v4
    for fi10 := 0; fi10 < 18; fi10++ {
        v9 = (float32((0.0 / 1.5)) * -2.25)
        v1 = append(v1, "hello")
        _ = fi10
    }
    _ = len(v1)
    if v4 {
        var v11 float32 = -2.25
        _ = v11
    } else {
        fmt.Println(v9, v9)
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        if (65535 == 664) {
            if v4 {
                v3 = v3
                switch "a" {
                case "\n\t":
                    if false {
                        var v13 bool = false
                        _ = v13
                    } else {
                        v3 -= (v3 - 2147483647)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                default:
                    var v14 string = ("hello" + ("" + "hello"))
                    _ = v14
                }
                var v15 uint = 65535
                _ = v15
            } else {
                v4 = (false || true)
            }
            switch ("" + ("\n\t" + "hello")) {
            case "":
                var v16 bool = v4
                _ = v16
            case "a":
                var v17 int8 = -1
                _ = v17
            case "foo bar":
                if v4 {
                    var v18 string = ("foo bar" + "hello")
                    _ = v18
                } else {
                    v3 -= (v3 / ((-870) | 1))
                }
            }
        } else {
            for fi19 := 0; fi19 < 8; fi19++ {
                v3 = v3
                _ = fi19
            }
            if v4 {
                if v4 {
                    v5[(("hello" + "") + "")] = fi12
                }
            }
        }
        _ = fi12
    }
    delete(v5, ("\n\t" + "foo bar"))
    if (v4 || (4294967295 >= 1)) {
        for fi20 := 0; fi20 < 20; fi20++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v21 := v2["a"]
            v21 -= (fi20 & 32768)
            _ = fi20
            _ = v21
        }
        if true {
            v9 = float32(0.1)
        }
    } else {
        _ = len(v1)
        _ = len(v5)
    }
    v4 = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch -711 {
    case -970:
        v22 := v5[""]
        _ = v22
    default:
        for fi23 := 0; fi23 < 13; fi23++ {
            v24 := GenS1{F0: 105}
            var v25 string = "foo bar"
            v26 := GenS1{F0: -128}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    }
    v9 = ((float32(((3.14159 + (-2.25 * 3.14159)) / 1.0)) + (v9 - float32(1e10))) + -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
