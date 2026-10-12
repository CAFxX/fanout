package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int, p1 uint16, p2 uint32) int32 {
    for fi4 := 0; fi4 < 10; fi4++ {
        p1 = p1
        _ = fi4
    }
    if false {
        if true {
            for fi5 := 0; fi5 < 4; fi5++ {
                var v6 int8 = -128
                _ = fi5
                _ = v6
            }
        }
        if false {
            var v7 string = ("\n\t" + "a")
            if true {
                v8 := GenS1{F0: (p2 / (((p2 | p2)) | 1))}
                _ = v8
            }
            _ = v7
        } else {
            switch "hello" {
            case "a":
                var v9 float32 = ((0.0 + 1e10) - ((1.5 + float32(((1e10 / (-2.25 - 0.1)) - 1.5))) - (-2.25 + 0.1)))
                _ = v9
            default:
                for fi10 := 0; fi10 < 19; fi10++ {
                    v11 := GenS1{F0: 65535}
                    _ = fi10
                    _ = v11
                }
            }
            p1 += 2
        }
    } else {
        p1 = p1
        for fi12 := 0; fi12 < 13; fi12++ {
            p1 = (p1 | uint16(p2))
            var v13 uint32 = (p2 + (p2 + (p2 + 1)))
            var v14 uint16 = (p1 >> ((42) & 63))
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    v15 := []int{128, -812, p0}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 237
}

func main() {
    var v1 int = 42
    var v2 uint16 = 0
    v3 := make(map[uint32]int)
    var v16 uint32 = 4294967295
    v17 := GenS1{F0: v16}
    fmt.Println(v1)
    var v18 uint32 = v16
    if (((false && true) && true) || false) {
        v18 = (v16 / ((1) | 1))
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 bool = ("hello" > ("" + ("hello" + "hello")))
            _ = fi19
            _ = v20
        }
        if true {
            v2 = v2
            fmt.Println(v16)
        }
    }
    v21 := []int{(v1 >> ((65535) & 63)), v1, (v1 & 2147483647), v1}
    var v22 uint8 = 255
    fmt.Println(v22)
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := GenS1{F0: uint32(v22)}
        _ = fi23
        _ = v24
    }
    switch (v1 ^ v1) {
    case -9223372036854775808:
        v22 = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
