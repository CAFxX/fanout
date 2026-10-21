package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0() uint {
    v4 := GenS1{F0: 4294967295}
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := []int{fi5, (fi5 | fi5), (^(fi5)), -629}
        _ = fi5
        _ = v6
    }
    v7 := []int{-86}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v8 := GenS1{F0: 702}
    _ = v4
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 bool = false
    var v2 float64 = 1.5
    v3 := make(map[int]string)
    var v9 bool = v1
    v2 = 1.5
    var v10 int32 = -1
    v3[32767] = ""
    v10 += v10
    v11 := GenS1{F0: 4294967295}
    for fi12 := 0; fi12 < 19; fi12++ {
        switch fi12 {
        case 2147483647:
            var v13 int = -2147483648
            _ = v13
        case 32767:
            delete(v3, fi12)
        }
        v14 := GenS2{F0: 0, F1: 42}
        if v1 {
            var v15 float64 = v2
            v1 = v1
            _ = v15
        }
        _ = fi12
        _ = v14
    }
    if (454 != 42) {
        for fi16 := 0; fi16 < 12; fi16++ {
            delete(v3, 0)
            for fi17 := 0; fi17 < 14; fi17++ {
                fmt.Println(fi17)
                _ = fi17
            }
            v18 := GenS3{F0: 42}
            _ = fi16
            _ = v18
        }
        var v19 int = 669
        _ = v19
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(v9)
            _ = fi20
        }
        switch 331 {
        case 32768:
            fmt.Println(v9, v2, v10)
        case -487:
            fmt.Println(v10, v9, v2)
        case 689:
            for fi21 := 0; fi21 < 5; fi21++ {
                fmt.Println(v1, v10)
                _ = fi21
            }
        default:
            v22 := GenS2{F0: 2, F1: 2}
            _ = v22
        }
        v10 -= v10
    }
    v10 -= 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
