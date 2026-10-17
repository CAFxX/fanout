package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint64) string {
    fmt.Println(p1)
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{fi7, (-(fi7)), (fi7 % (((fi7 & 9223372036854775807)) | 1))}
        p1 = 329
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return ("a" + "hello")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 65535
    v3 := make(map[uint32]int)
    var v4 string = "\n\t"
    var v5 int16 = 0
    v6 := make(map[int]uint32)
    var v9 uint = 0
    fmt.Println(v2, v5)
    switch "a" {
    case "hello":
        if true {
            v10 := GenS1{F0: 0, F1: (^(v2))}
            _ = v10
        }
    }
    var v11 string = (v4 + v4)
    v12 := GenS1{F0: 32768, F1: (v2 / ((v2) | 1))}
    switch -93 {
    case 42:
        fmt.Println(v9, v11)
    case -542:
        v13 := GenS2{F0: 758}
        _ = v13
    }
    fmt.Println(v11)
    v6[0] = 0
    v5 += v5
    var v14 int8 = 65
    for fi15 := 0; fi15 < 5; fi15++ {
        v14 += v14
        _ = fi15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        switch "" {
        case "\n\t":
            fmt.Println(v4)
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            v4 = v11
            _ = fi17
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            if ((v9 & (v9 * uint(v2))) >= 1) {
                var v19 uint = (v9 >> ((v9) & 63))
                v20 := []int{fi18, fi16}
                _ = v19
                _ = v20
            } else {
                v2 -= v2
                var v21 bool = (!true)
                _ = v21
            }
            _ = fi18
        }
        _ = fi16
    }
    if true {
        v9 += 2
    }
    var v22 int64 = 1
    v2 = (v2 % (((v2 & (v2 % ((v2) | 1)))) | 1))
    if false {
        fmt.Println(v5)
        v5 -= -1
        fmt.Println(v14, v5, v11)
    }
    var v23 string = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
