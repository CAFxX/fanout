package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint8; F1 uint32; F2 int16 }
func genFn0(p0 string, p1 string, p2 int32) int16 {
    var v7 int8 = 1
    v8 := []int{-9223372036854775808, 32767, -625}
    v7 -= v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 float64 = float64(0.1)
    var v2 int8 = 0
    var v3 bool = (18446744073709551615 != 519)
    var v4 int16 = 1000
    v5 := make(map[uint32]string)
    v6 := make(map[int]uint32)
    v9 := GenS1{F0: 784, F1: true}
    v1 = v1
    if (42 < 127) {
        for fi10 := 0; fi10 < 3; fi10++ {
            delete(v6, fi10)
            _ = fi10
        }
        var v11 int8 = v2
        _ = v11
    }
    v3 = false
    v1 -= 0.1
    switch 4294967295 {
    case -9223372036854775808:
        switch 9223372036854775807 {
        case 1:
            v2 = -1
        }
    case -2147483648:
        var v12 float32 = (0.0 - 1.5)
        _ = v12
    }
    v13 := GenS1{F0: 0, F1: false}
    if false {
        v14 := GenS2{F0: true}
        fmt.Println(v3)
        if ((0.1 == 1.5) || true) {
            v15 := []int{}
            fmt.Println(v3)
            switch -9223372036854775808 {
            case -105:
                v16 := GenS3{F0: -9223372036854775808}
                _ = v16
            case 2147483647:
                var v17 bool = true
                _ = v17
            }
            _ = v15
        }
        _ = v14
    } else {
        v5[127] = "\n\t"
        v3 = (255 == 4294967295)
        v2 -= -27
    }
    v3 = true
    _ = len(v6)
    v18 := GenS4{F0: v2, F1: (("a" + "\n\t") + "hello")}
    for fi19 := 0; fi19 < 18; fi19++ {
        if true {
            v20 := GenS1{F0: 0, F1: (722 <= 1)}
            v21 := GenS4{F0: (v2 | v2), F1: "hello"}
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 uint64 = 0
                _ = fi22
                _ = v23
            }
            _ = v20
            _ = v21
        } else {
            v24 := []int{(fi19 - fi19), fi19, (-(fi19))}
            v2 = (v2 | (v2 - (v2 + v2)))
            var v25 float32 = 0.1
            _ = v24
            _ = v25
        }
        v26 := []int{}
        switch int(v2) {
        case -903:
            v27 := GenS2{F0: ("a" >= ("foo bar" + ("foo bar" + ("foo bar" + "a"))))}
            _ = v27
        }
        _ = fi19
        _ = v26
    }
    v3 = v3
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 int = fi28
        v4 = (v4 << ((81) & 63))
        v1 = 0.1
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        var v31 uint16 = 974
        for fi32 := 0; fi32 < 10; fi32++ {
            v33 := []int{fi30, (-(fi32))}
            var v34 bool = v3
            v33 = append(v33, fi30)
            _ = fi32
            _ = v33
            _ = v34
        }
        v2 += (v2 / ((v2) | 1))
        _ = fi30
        _ = v31
    }
    if false {
        for fi35 := 0; fi35 < 9; fi35++ {
            switch "a" {
            case "\n\t":
                var v36 uint8 = 1
                _ = v36
            case "foo bar":
                var v37 uint8 = 2
                _ = v37
            }
            v38 := GenS3{F0: (fi35 * int(v4))}
            v4 = v4
            _ = fi35
            _ = v38
        }
        v39 := GenS5{F0: 10, F1: 2147483647, F2: -620}
        _ = v39
    }
    v4 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
