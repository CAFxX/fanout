package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 string) uint64 {
    var v4 uint8 = 112
    p1 = ("" + (p1 + "\n\t"))
    var v5 int32 = -704
    if false {
        var v6 uint32 = 2147483648
        var v7 int32 = (-(v5))
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 string = ("foo bar" + "")
    v2 := make(map[int]string)
    v3 := []string{v1, v1, ""}
    if ((v1 < v1) && false) {
        if false {
            switch 32767 {
            case -276:
                v1 = ("hello" + "")
            case 32768:
                fmt.Println(v1)
            case 0:
                fmt.Println(v1)
            }
        } else {
            _ = len(v2)
            v1 = v1
            if true {
                if false {
                    fmt.Println(v1)
                    if (32767 > 1) {
                        var v8 bool = false
                        _ = v8
                    }
                }
                if len(v3) > 0 {
                    v3[0] = (v1 + "hello")
                }
            } else {
                switch "a" {
                case "foo bar":
                    v1 = "a"
                case "a":
                    var v9 uint32 = 1
                    _ = v9
                }
                var v10 int16 = -169
                var v11 int32 = -885
                _ = v10
                _ = v11
            }
        }
    } else {
        var v12 int16 = 15
        var v13 uint64 = 1
        _ = v12
        _ = v13
    }
    v1 = ("a" + v1)
    v14 := GenS1{F0: 255}
    if true {
        var v15 string = "\n\t"
        fmt.Println(v15)
        if len(v3) > 0 {
            v3[0] = ""
        }
        _ = v15
    }
    var v16 string = (v1 + v1)
    switch v1 {
    case "hello":
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := GenS2{F0: (v1 + (v16 + (v1 + v1)))}
            var v19 float32 = 0.1
            for fi20 := 0; fi20 < 4; fi20++ {
                v21 := v2[2147483647]
                _ = fi20
                _ = v21
            }
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    v3 = append(v3, "hello")
    _ = len(v3)
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int64 = 42
        _ = fi22
        _ = v23
    }
    var v24 int = 9223372036854775807
    v25 := v2[(v24 + v24)]
    v1 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
