package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := []int{(fi4 << ((0) & 63)), fi4, (fi4 * fi4)}
        var v6 uint32 = 86
        _ = fi4
        _ = v5
        _ = v6
    }
    switch -912 {
    case 274:
        if (true || (false && false)) {
            var v7 int8 = -1
            v7 = v7
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint8 = 2
        if (42 != 42) {
            var v10 uint8 = v9
            if ((0 >= 958) && false) {
                var v11 float32 = 1.5
                _ = v11
            } else {
                v12 := []int{(fi8 / ((32768) | 1))}
                var v13 bool = true
                _ = v12
                _ = v13
            }
            v10 = v10
            _ = v10
        }
        if (!(0 >= 425)) {
            switch "\n\t" {
            case "foo bar":
                var v14 uint64 = 0
                _ = v14
            case "a":
                var v15 float64 = (-2.25 + ((float64((3.14159 * 3.14159)) / -2.25) / 1.5))
                _ = v15
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                v17 := []int{fi16, 0, -740}
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 6; fi18++ {
                v19 := []int{}
                switch "\n\t" {
                case "hello":
                    var v20 uint32 = 42
                    _ = v20
                case "a":
                    v9 = 182
                case "":
                    v9 += (^(v9))
                default:
                    _ = len(v19)
                }
                _ = fi18
                _ = v19
            }
        } else {
            v21 := GenS1{F0: true, F1: 1}
            v22 := GenS2{F0: 32767}
            if true {
                fmt.Println(fi8)
                var v23 uint = 42
                v9 += v9
                _ = v23
            }
            _ = v21
            _ = v22
        }
        _ = fi8
        _ = v9
    }
    var v24 int32 = 42
    for fi25 := 0; fi25 < 15; fi25++ {
        v24 = -2147483648
        var v26 int8 = 100
        _ = fi25
        _ = v26
    }
    _ = v24
    return 2147483647
}

func main() {
    v1 := []int{101, 128, 127, 280}
    var v2 uint = 255
    var v3 int64 = 0
    if (18446744073709551615 > 42) {
        v1 = append(v1, 42)
        v27 := GenS1{F0: (!(!false)), F1: 2147483648}
        fmt.Println(v2, v3)
        _ = v27
    }
    if ("\n\t" < "foo bar") {
        v2 = (v2 % ((v2) | 1))
        if true {
            var v28 uint8 = 2
            if true {
                fmt.Println(v3, v3)
                var v29 bool = true
                v30 := GenS2{F0: -485}
                _ = v29
                _ = v30
            } else {
                v28 += (^(v28))
                var v31 string = "a"
                _ = v31
            }
            v28 -= (v28 % (((v28 + v28)) | 1))
            _ = v28
        }
        fmt.Println(v3)
    }
    fmt.Println(v3, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 int32 = -616
        _ = fi32
        _ = v33
    }
    v2 += v2
    v34 := GenS2{F0: 0}
    fmt.Println(v3)
    var v35 float64 = 1e10
    switch ("" + ("a" + "\n\t")) {
    case "foo bar":
        v2 = v2
    default:
        v3 += (v3 >> (((v2 + v2)) & 63))
    }
    var v36 int32 = 458
    if len(v1) > 0 {
        v1[0] = 728
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
