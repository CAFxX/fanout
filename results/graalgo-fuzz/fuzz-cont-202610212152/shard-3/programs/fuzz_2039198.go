package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 string }
type GenS2 struct { F0 string; F1 uint32; F2 uint8 }
type GenS3 struct { F0 string; F1 int; F2 uint16 }
func genFn0(p0 bool, p1 uint16, p2 bool) uint8 {
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(p0, p1, p0)
        v7 := []int{fi6}
        fmt.Println(p1, p2, p0)
        _ = fi6
        _ = v7
    }
    var v8 uint16 = (p1 << ((2) & 63))
    switch 2147483647 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 5; fi9++ {
            switch (fi9 | 0) {
            case -9223372036854775808:
                var v10 bool = (!(1 <= 42))
                _ = v10
            case 0:
                for fi11 := 0; fi11 < 9; fi11++ {
                    v12 := GenS1{F0: fi9, F1: 95, F2: ("a" + "")}
                    v13 := []int{(fi9 ^ (fi9 | (^(fi11))))}
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 42:
                var v14 int = (fi9 | fi9)
                _ = v14
            }
            _ = fi9
        }
    case 510:
        switch "foo bar" {
        case "":
            fmt.Println(p0, p0)
        }
    }
    var v15 bool = (true || p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v8
    return uint8(v8)
}

func main() {
    var v1 uint = 255
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    var v5 int = 967
    var v16 uint8 = 131
    if (-1 > 1) {
        var v17 uint32 = 2
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint32 = v17
            fmt.Println(v19, v1, v16)
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        fmt.Println(v1)
        _ = fi20
    }
    switch ("a" + "a") {
    case "foo bar":
        v1 -= v1
    case "\n\t":
        v21 := GenS1{F0: (v5 % ((32768) | 1)), F1: v16, F2: "hello"}
        _ = v21
    }
    v5 = v5
    var v22 int64 = int64(v16)
    v23 := GenS2{F0: ("hello" + "foo bar"), F1: 2, F2: genFn0(false, 65535, (2 != 232))}
    v1 += v1
    var v24 float64 = -2.25
    for fi25 := 0; fi25 < 3; fi25++ {
        v5 = (v5 ^ 32768)
        v26 := GenS3{F0: "", F1: fi25, F2: 786}
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
