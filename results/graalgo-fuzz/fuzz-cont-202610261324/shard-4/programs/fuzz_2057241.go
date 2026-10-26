package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 bool = false
    v2 := []int{32767, 428}
    var v3 int32 = 2147483647
    v1 = true
    fmt.Println(v1, v3)
    fmt.Println(v3)
    switch "hello" {
    case "a":
        if (true && (v1 && false)) {
            fmt.Println(v3, v3)
            if len(v2) > 0 {
                v2[0] = 0
            }
        } else {
            _ = len(v2)
        }
    case "foo bar":
        v3 -= v3
    }
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 float32 = float32((3.14159 * 1e10))
        _ = fi4
        _ = v5
    }
    if v1 {
        var v6 int16 = 0
        if true {
            v1 = (65535 != 1)
            var v7 uint64 = 0
            _ = v7
        } else {
            var v8 int16 = v6
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: v3, F1: 520, F2: 0}
                if false {
                    v2 = append(v2, (^(fi9)))
                    var v11 uint = 1
                    if v1 {
                        var v12 int32 = v3
                        v2 = append(v2, (fi9 | (fi9 - fi9)))
                        _ = v12
                    }
                    _ = v11
                }
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 int = (fi13 - fi9)
                    v15 := GenS1{F0: v3, F1: 42, F2: 127}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi9
                _ = v10
            }
            var v16 uint = 2
            _ = v16
            _ = v8
        }
        _ = v6
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        _ = len(v2)
        if true {
            v1 = ((v3 | (v3 ^ (v3 + -2147483648))) <= int32(fi17))
            var v18 float64 = (3.14159 * -2.25)
            _ = v18
        }
        if (v1 || false) {
            for fi19 := 0; fi19 < 9; fi19++ {
                v1 = (!true)
                v20 := GenS2{F0: 42}
                var v21 string = (((("foo bar" + ((("foo bar" + ("" + "hello")) + "hello") + "")) + "hello") + "") + "\n\t")
                _ = fi19
                _ = v20
                _ = v21
            }
        } else {
            v2 = append(v2, 293)
        }
        _ = fi17
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v1 = true
        _ = fi22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        v3 = (v3 << ((0) & 63))
        var v24 string = "hello"
        _ = fi23
        _ = v24
    }
    var v25 float32 = 0.1
    v25 = 1e10
    _ = len(v2)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
}
