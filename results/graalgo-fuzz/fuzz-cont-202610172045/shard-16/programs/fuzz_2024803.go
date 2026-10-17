package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int64 }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint8 }
func genFn0() int {
    var v7 uint32 = 2147483647
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := []int{(-(fi9)), fi9, -1}
            if (42 > 233) {
                v7 = (v7 % ((v7) | 1))
            } else {
                v11 := GenS1{F0: 1, F1: 3, F2: 605}
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if (((1 != 127) && false) || ((3.14159 < 1.5) || false)) {
        fmt.Println(v7)
        var v12 uint16 = 255
        v12 = 650
        _ = v12
    }
    var v13 uint64 = 1
    v14 := GenS1{F0: uint32(v13), F1: 1, F2: 9223372036854775807}
    _ = v13
    _ = v14
    _ = v7
    return 128
}
func genFn1() uint64 {
    var v15 uint64 = 4294967295
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int64 = -9223372036854775808
        v18 := GenS2{F0: 2147483647, F1: v17, F2: 1}
        fmt.Println(v17)
        _ = fi16
        _ = v17
        _ = v18
    }
    if true {
        var v19 string = (("hello" + ("" + "foo bar")) + ("hello" + "a"))
        var v20 int64 = 413
        _ = v19
        _ = v20
    }
    _ = v15
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]uint32)
    var v3 float64 = 1.5
    var v4 int64 = int64(v1)
    var v5 float64 = v3
    var v6 float32 = float32(v3)
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := []int{-668, (fi21 >> ((65535) & 63)), (fi21 | fi21)}
        fmt.Println(v5)
        v6 += 1.5
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            if false {
                v4 = v4
                delete(v2, fi24)
            }
            v25 := []int{fi24, (-(fi24)), fi23, fi24}
            switch 647 {
            case -726:
                switch (^(fi23)) {
                case -165:
                    var v26 uint8 = 128
                    _ = v26
                case 205:
                    if false {
                        v27 := GenS3{F0: 42}
                        v1 += v1
                        var v28 int64 = -414
                        _ = v27
                        _ = v28
                    } else {
                        var v29 int64 = (v4 >> ((1) & 63))
                        var v30 float64 = 1e10
                        _ = v29
                        _ = v30
                    }
                }
            case 753:
                if (v1 <= (^(v1))) {
                    var v31 uint8 = 0
                    _ = v31
                } else {
                    v4 = 1
                    switch (fi23 + genFn0()) {
                    case -325:
                        v32 := GenS1{F0: 2, F1: int8(v4), F2: 42}
                        _ = v32
                    case -2147483648:
                        v6 = (float32(v5) / -2.25)
                    }
                    fmt.Println(fi23)
                }
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                v33 := GenS4{F0: true, F1: -2147483648}
                _ = v33
            }
            _ = fi24
            _ = v25
        }
        for fi34 := 0; fi34 < 5; fi34++ {
            var v35 int8 = 127
            v36 := GenS5{F0: 1}
            if true {
                v3 -= (v5 + (0.0 / -2.25))
                var v37 uint64 = 1
                _ = v37
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        for fi38 := 0; fi38 < 5; fi38++ {
            v5 = v3
            _ = fi38
        }
        _ = fi23
    }
    v1 = (v1 << ((18446744073709551615) & 63))
    v39 := GenS3{F0: 18446744073709551615}
    fmt.Println(v4)
    var v40 int16 = 53
    v41 := GenS5{F0: 128}
    v42 := []int{-1, 127}
    delete(v2, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
