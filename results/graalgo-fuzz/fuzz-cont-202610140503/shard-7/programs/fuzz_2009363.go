package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 float32 = 1e10
    v2 := []int{32768}
    v3 := make(map[int]string)
    v4 := []string{}
    var v5 uint32 = 42
    if (!true) {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: false}
            var v8 uint32 = 2
            switch ("foo bar" + "") {
            case "":
                switch 9223372036854775807 {
                case 670:
                    v5 = v8
                case -9223372036854775808:
                    var v9 int64 = -600
                    _ = v9
                case 505:
                    if false {
                        _ = len(v3)
                        _ = len(v3)
                    } else {
                        for fi10 := 0; fi10 < 14; fi10++ {
                            fmt.Println(v8)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi10
                        }
                    }
                }
            case "a":
                v11 := GenS2{F0: "hello"}
                _ = v11
            case "foo bar":
                var v12 bool = true
                _ = v12
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v5 = v5
        _ = fi13
    }
    var v14 int = -2147483648
    var v15 int8 = 100
    var v16 uint16 = 528
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v17 := v3[(v14 - v14)]
    var v18 uint16 = (v16 * 32767)
    switch "" {
    case "a":
        if (-1 == -1) {
            v19 := GenS3{F0: -787}
            _ = v19
        }
    default:
        v14 = 32767
    }
    if true {
        v20 := GenS4{F0: v5}
        var v21 uint16 = (v18 & v18)
        _ = v20
        _ = v21
    }
    if (v17 < v17) {
        var v22 uint16 = 42
        if false {
            v17 = "hello"
            v14 = v14
            var v23 uint = 1
            _ = v23
        }
        _ = v22
    }
    delete(v3, 605)
    fmt.Println(v17)
    var v24 uint32 = (v5 << ((2) & 63))
    v2 = append(v2, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
