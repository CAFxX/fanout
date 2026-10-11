package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int64 }
func genFn0(p0 uint, p1 int16, p2 uint16) string {
    p1 = (p1 + 32767)
    v5 := []int{128, -450}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "a"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 0
    var v3 uint16 = (^(v2))
    v4 := make(map[int]string)
    if true {
        v6 := GenS1{F0: 1, F1: 0, F2: false}
        _ = v6
    }
    var v7 float64 = (((float64(3.14159) - 1e10) - (3.14159 - -2.25)) / 1.5)
    if false {
        v3 += v3
        switch -1 {
        case 238:
            v2 += (^(v2))
        case -2147483648:
            switch ("a" + ("hello" + "foo bar")) {
            case "a":
                fmt.Println(v3, v2)
            }
        case 127:
            var v8 int8 = 127
            _ = v8
        }
    }
    var v9 float64 = (((3.14159 - 3.14159) * v7) / v7)
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(v9, v3)
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    v12 := v1[4294967295]
    switch 0 {
    case 507:
        fmt.Println(v2)
    case 128:
        var v13 bool = false
        _ = v13
    }
    if (-288 <= 533) {
        v9 = v7
    }
    if (42 >= 160) {
        if true {
            var v14 uint8 = uint8(v3)
            v14 = 128
            _ = v14
        }
        var v15 float64 = v9
        _ = v15
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            switch fi16 {
            case -294:
                v3 = v3
            case 4294967295:
                v17 := GenS1{F0: 1, F1: 32767, F2: (978 <= 2)}
                _ = v17
            case 0:
                var v18 float64 = v9
                _ = v18
            }
            var v19 float32 = -2.25
            v20 := GenS2{F0: -547}
            _ = fi16
            _ = v19
            _ = v20
        }
        v21 := GenS2{F0: -71}
        var v22 int8 = 127
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: 1, F1: -9223372036854775808}
    if v12 {
        if false {
            for fi24 := 0; fi24 < 18; fi24++ {
                var v25 string = ""
                _ = fi24
                _ = v25
            }
            var v26 bool = (((v7 / v9) >= v9) || false)
            _ = v26
        }
    } else {
        var v27 int8 = 1
        switch "\n\t" {
        case "\n\t":
            var v28 int64 = 0
            _ = v28
        case "":
            v29 := GenS1{F0: 2147483647, F1: 42, F2: ((v3 / ((v2) | 1)) >= v2)}
            _ = v29
        case "a":
            var v30 int64 = 1
            _ = v30
        }
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
