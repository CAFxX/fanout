package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 string; F2 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 34
    var v3 string = "\n\t"
    var v4 uint = 473
    var v5 int16 = 32767
    var v6 int16 = (v5 >> ((v4) & 63))
    var v7 uint64 = 0
    v8 := GenS1{F0: -128}
    v9 := GenS1{F0: 42}
    fmt.Println(v6, v4)
    if (0 == 15) {
        var v10 int16 = v6
        v7 += (v7 + 4294967295)
        v2 -= (v2 | (v2 << (((v4 + v4)) & 63)))
        _ = v10
    } else {
        if false {
            if (false || false) {
                var v11 uint = (v4 / ((0) | 1))
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := []int{}
                    _ = fi12
                    _ = v13
                }
                var v14 string = v3
                _ = v11
                _ = v14
            } else {
                var v15 int = v2
                v6 -= 0
                v15 = v15
                _ = v15
            }
            fmt.Println(v4, v4, v6)
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := []int{fi16, (fi16 * (fi16 + (-(fi16)))), (fi16 % (((v2 | v2)) | 1))}
                var v18 uint16 = 799
                v19 := GenS1{F0: -73}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        } else {
            fmt.Println(v2)
        }
        var v20 float32 = 1e10
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v2 = fi21
        _ = fi21
    }
    if ((108 != 41) || false) {
        switch 127 {
        case -2147483648:
            var v22 uint8 = 2
            _ = v22
        case -129:
            v6 = (^(v6))
        case 678:
            v23 := GenS1{F0: 42}
            _ = v23
        }
        _ = len(v1)
        var v24 string = v3
        _ = v24
    }
    switch v2 {
    case 23:
        for fi25 := 0; fi25 < 18; fi25++ {
            if ("" < v3) {
                v26 := GenS2{F0: 1, F1: "hello", F2: 0}
                for fi27 := 0; fi27 < 15; fi27++ {
                    v28 := GenS1{F0: -22}
                    fmt.Println(v4)
                    var v29 uint32 = 530
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            } else {
                v30 := GenS1{F0: 99}
                _ = v30
            }
            var v31 uint16 = 878
            v32 := []int{}
            _ = fi25
            _ = v31
            _ = v32
        }
    case -471:
        if (128 != 1) {
            fmt.Println(v2, v6)
            if (187 >= 127) {
                v5 += 1
                var v33 int = (v2 - 36)
                v34 := GenS1{F0: 20}
                _ = v33
                _ = v34
            } else {
                v35 := []int{}
                var v36 uint = (v4 - (v4 / ((v4) | 1)))
                _ = v35
                _ = v36
            }
        }
    case -9223372036854775808:
        v37 := GenS1{F0: -128}
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
