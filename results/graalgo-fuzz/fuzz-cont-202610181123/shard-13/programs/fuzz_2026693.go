package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 string; F2 int64 }

func main() {
    v1 := []string{"foo bar", ""}
    var v2 int = 127
    v3 := make(map[string]int)
    var v4 int16 = 224
    for fi5 := 0; fi5 < 18; fi5++ {
        if false {
            fmt.Println(v2)
        } else {
            for fi6 := 0; fi6 < 19; fi6++ {
                if (!(3.14159 >= 0.1)) {
                    fmt.Println(fi6)
                } else {
                    v1 = append(v1, "a")
                }
                switch ("a" + "hello") {
                case "foo bar":
                    v4 -= -1
                    fmt.Println(fi5, v4, fi6)
                case "":
                    v7 := GenS1{F0: 255}
                    _ = v7
                case "a":
                    v2 = fi6
                default:
                    v8 := GenS1{F0: 0}
                    _ = v8
                }
                if (true && (127 >= 42)) {
                    var v9 uint64 = 42
                    switch ((("a" + (("" + ("" + "foo bar")) + "")) + ("foo bar" + "foo bar")) + ("hello" + ("\n\t" + ("foo bar" + "\n\t")))) {
                    case "":
                        v1 = append(v1, "a")
                    case "a":
                        var v10 int = v2
                        _ = v10
                    case "hello":
                        var v11 uint = 1
                        _ = v11
                    default:
                        if false {
                            v12 := GenS1{F0: 0}
                            v9 = v9
                            fmt.Println(v2, v2)
                            _ = v12
                        } else {
                            v13 := GenS1{F0: 255}
                            v14 := GenS1{F0: 128}
                            _ = v13
                            _ = v14
                        }
                    }
                    _ = v9
                }
                _ = fi6
            }
            fmt.Println(v4, fi5, v4)
            v4 = v4
        }
        var v15 uint32 = 65535
        _ = fi5
        _ = v15
    }
    v3[("hello" + "a")] = (v2 % ((-568) | 1))
    fmt.Println(v4, v2)
    delete(v3, "a")
    var v16 string = ("foo bar" + "foo bar")
    var v17 string = v16
    fmt.Println(v2)
    if (0 != -1) {
        v16 = v16
        v16 = "foo bar"
    }
    fmt.Println(v17, v17)
    var v18 uint16 = 65535
    var v19 float64 = 1.5
    if (!true) {
        switch v17 {
        case "foo bar":
            if (!(0.0 < float64(0.1))) {
                v18 = 32767
                fmt.Println(v4, v19)
            }
        default:
            if (!true) {
                var v20 uint16 = v18
                v21 := v3[""]
                v22 := GenS2{F0: -465, F1: v16, F2: 42}
                _ = v20
                _ = v21
                _ = v22
            } else {
                var v23 int16 = (v4 >> ((0) & 63))
                _ = v23
            }
        }
        v24 := GenS1{F0: 127}
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        fmt.Println(fi25, fi25)
        _ = fi25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi26
    }
    var v27 int16 = v4
    v28 := GenS1{F0: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
