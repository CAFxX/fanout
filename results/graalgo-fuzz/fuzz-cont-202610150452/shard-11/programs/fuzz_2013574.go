package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0(p0 bool, p1 int) uint8 {
    for fi7 := 0; fi7 < 4; fi7++ {
        if p0 {
            switch ("hello" + "foo bar") {
            case "foo bar":
                for fi8 := 0; fi8 < 4; fi8++ {
                    v9 := []int{}
                    if true {
                        if (p0 || (-707 > 42)) {
                            var v10 int32 = -1
                            _ = len(v9)
                            p1 += fi8
                            _ = v10
                        }
                        p1 -= fi8
                        var v11 uint16 = 1
                        _ = v11
                    } else {
                        fmt.Println(p0, fi7)
                        p1 += -344
                    }
                    _ = fi8
                    _ = v9
                }
            }
            var v12 string = "\n\t"
            _ = v12
        }
        v13 := []int{fi7}
        var v14 bool = false
        _ = fi7
        _ = v13
        _ = v14
    }
    p1 -= p1
    var v15 int32 = 0
    switch ("a" + "hello") {
    case "foo bar":
        v16 := GenS1{F0: 563, F1: 255, F2: 1}
        _ = v16
    case "":
        if true {
            switch "hello" {
            case "a":
                v15 -= (v15 / ((v15) | 1))
            case "hello":
                v17 := GenS1{F0: 32767, F1: 42, F2: 969}
                _ = v17
            case "\n\t":
                v15 -= v15
            }
            fmt.Println(p0)
        } else {
            v15 -= -728
        }
    }
    _ = p0
    _ = p1
    _ = v15
    return 49
}
func genFn1(p0 int16, p1 int16) int32 {
    if ((false && false) || true) {
        v18 := GenS2{F0: 127, F1: false}
        _ = v18
    } else {
        p0 = 32767
        for fi19 := 0; fi19 < 8; fi19++ {
            if (2 > 127) {
                p0 += p0
                v20 := GenS3{F0: 4294967295, F1: (^(fi19))}
                p1 = 1
                _ = v20
            }
            v21 := []int{fi19, (fi19 << ((255) & 63)), (fi19 | 0), 341}
            _ = fi19
            _ = v21
        }
        v22 := GenS3{F0: 0, F1: 32768}
        _ = v22
    }
    p1 = (p1 * p1)
    if false {
        var v23 string = "a"
        var v24 int64 = -9223372036854775808
        _ = v23
        _ = v24
    }
    v25 := []int{32768, -2147483648, 32768, 619}
    _ = p0
    _ = p1
    _ = v25
    return 0
}

func main() {
    var v1 uint8 = 1
    var v2 string = "hello"
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v5 bool = (false || (65535 <= 363))
    v6 := make(map[int]bool)
    var v26 int32 = -2147483648
    if v5 {
        switch (("foo bar" + v2) + v2) {
        case "":
            if false {
                var v27 uint8 = 139
                _ = v27
            } else {
                v28 := []int{32768, -296, 536}
                if (!((3.14159 + 3.14159) <= 1e10)) {
                    if (!v5) {
                        v3[2147483647] = ((v2 + v2) + "hello")
                        _ = len(v28)
                        for fi29 := 0; fi29 < 9; fi29++ {
                            v1 += (v1 >> ((42) & 63))
                            v30 := GenS2{F0: 1, F1: (772 >= 2147483648)}
                            _ = fi29
                            _ = v30
                        }
                    }
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v32 := GenS2{F0: 35, F1: (v2 <= v2)}
                        _ = fi31
                        _ = v32
                    }
                    v2 = v2
                }
                _ = v28
            }
        case "foo bar":
            var v33 int16 = -1
            _ = v33
        case "hello":
            v34 := []int{-580, -455, -450}
            _ = v34
        }
    }
    v2 = "\n\t"
    v2 = "foo bar"
    if (!(true || false)) {
        if (v26 <= v26) {
            if v5 {
                var v35 string = v2
                var v36 uint16 = 32767
                v26 = (^(v26))
                _ = v35
                _ = v36
            } else {
                v26 = (v26 >> ((381) & 63))
            }
        }
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        if (65535 > uint(v1)) {
            v26 -= 1
            v1 = v1
            v38 := GenS4{F0: 2, F1: 32767}
            _ = v38
        }
        v39 := GenS1{F0: 792, F1: 0, F2: 2}
        v2 = v2
        _ = fi37
        _ = v39
    }
    v4[47] = v2
    var v40 float32 = float32(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
