package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
func genFn0(p0 int32, p1 string, p2 bool) uint8 {
    fmt.Println(p2, p1)
    var v5 float32 = 0.1
    v5 = v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}
func genFn1(p0 string) uint16 {
    v6 := GenS1{F0: 32768}
    switch -175 {
    case 32767:
        var v7 bool = false
        _ = v7
    case -1:
        p0 = p0
    case 4294967295:
        fmt.Println(p0)
    }
    v8 := GenS2{F0: p0, F1: false, F2: -68}
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            if false {
                var v11 int = 32768
                fmt.Println(v11)
                _ = v11
            }
            v12 := []int{fi9, 2147483647}
            fmt.Println(fi10, fi10, fi10)
            v13 := GenS1{F0: (fi9 ^ (fi9 * fi9))}
            switch (fi9 ^ 0) {
            case 9223372036854775807:
                p0 = (p0 + p0)
            case 0:
                v14 := GenS2{F0: ((("foo bar" + (p0 + "foo bar")) + (p0 + "\n\t")) + p0), F1: (0 != 0), F2: -1}
                _ = v14
            case 555:
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
            _ = fi10
            _ = v12
            _ = v13
        }
        var v15 uint16 = 2
        v16 := GenS1{F0: fi9}
        _ = fi9
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v6
    _ = v8
    return 0
}

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 2
    v3 := []string{"\n\t"}
    var v4 int64 = 9223372036854775807
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 uint64 = v2
        var v19 int8 = -120
        switch ("hello" + "foo bar") {
        case "":
            _ = len(v3)
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v1)
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int8 = 0
        _ = fi20
        _ = v21
    }
    var v22 int64 = (v4 >> ((755) & 63))
    var v23 float64 = ((3.14159 / (0.0 + 1e10)) + (0.1 * 1e10))
    v4 = -103
    fmt.Println(v1, v23, v22)
    var v24 int32 = 469
    var v25 uint16 = genFn1("hello")
    v26 := GenS2{F0: "", F1: true, F2: 113}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
