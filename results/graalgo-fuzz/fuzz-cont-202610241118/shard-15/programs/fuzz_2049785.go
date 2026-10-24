package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint8) int16 {
    v4 := []int{}
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    if p0 {
        fmt.Println(p1, p0)
        p0 = p0
    }
    v4 = append(v4, -149)
    _ = p0
    _ = p1
    _ = v4
    return -32768
}
func genFn1(p0 uint, p1 bool) bool {
    p0 = 65535
    p1 = p1
    v5 := []int{4294967295, -1}
    _ = p0
    _ = p1
    _ = v5
    return (255 < 42)
}

func main() {
    var v1 uint16 = 0
    var v2 uint16 = (v1 & v1)
    v3 := []string{}
    fmt.Println(v1, v2)
    var v6 uint64 = 647
    var v7 int = 1
    v3 = append(v3, "\n\t")
    var v8 int = v7
    v2 = (v2 ^ (^(v1)))
    fmt.Println(v1, v1)
    v9 := GenS1{F0: true, F1: 128, F2: (3.14159 != 0.0)}
    v1 = (v2 << ((640) & 63))
    v1 = 647
    switch v7 {
    case 0:
        var v10 string = "hello"
        _ = v10
    case 127:
        for fi11 := 0; fi11 < 19; fi11++ {
            v6 = (v6 ^ 0)
            _ = fi11
        }
    case -257:
        v12 := GenS1{F0: false, F1: 0, F2: (v2 > v2)}
        _ = v12
    }
    if (42 > -187) {
        var v13 uint = uint(v2)
        if true {
            var v14 int = 128
            if (!true) {
                v8 = v14
                for fi15 := 0; fi15 < 14; fi15++ {
                    v16 := GenS1{F0: (((^(v14)) <= fi15) || false), F1: 178, F2: true}
                    _ = fi15
                    _ = v16
                }
            } else {
                var v17 uint32 = 2147483648
                _ = v17
            }
            v14 -= v8
            _ = v14
        }
        if (v8 < v8) {
            var v18 int8 = -81
            _ = v18
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                switch "\n\t" {
                case "":
                    var v20 int8 = 100
                    _ = v20
                case "hello":
                    for fi21 := 0; fi21 < 20; fi21++ {
                        var v22 int32 = 0
                        _ = fi21
                        _ = v22
                    }
                }
                _ = fi19
            }
            var v23 string = "foo bar"
            fmt.Println(v13, v23)
            _ = v23
        }
        _ = v13
    } else {
        for fi24 := 0; fi24 < 3; fi24++ {
            var v25 int = 0
            v26 := GenS1{F0: true, F1: 255, F2: true}
            for fi27 := 0; fi27 < 8; fi27++ {
                fmt.Println(v8, v25)
                if true {
                    switch ("\n\t" + ("\n\t" + ("a" + "a"))) {
                    case "a":
                        fmt.Println(v8, v6)
                    }
                    var v28 int32 = 1
                    _ = v28
                }
                _ = fi27
            }
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    v29 := GenS1{F0: false, F1: 255, F2: (!false)}
    v30 := GenS2{F0: 45}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
