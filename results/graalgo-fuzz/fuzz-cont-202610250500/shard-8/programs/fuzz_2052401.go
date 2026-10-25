package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint }
type GenS3 struct { F0 int8 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := []int{}
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 uint = 257
            fmt.Println(fi6)
            switch int(v7) {
            case 115:
                v7 -= 18446744073709551615
            case 2147483647:
                v7 = v7
            }
            if (1 != 32767) {
                v8 := GenS1{F0: v7, F1: 42}
                switch fi4 {
                case 270:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                if (845 == uint16(v7)) {
                    v9 := GenS2{F0: 42, F1: -1, F2: v7}
                    _ = v9
                } else {
                    var v10 uint32 = 601
                    _ = v10
                }
                _ = v8
            }
            switch ("\n\t" + ("a" + "")) {
            case "hello":
                var v11 int32 = 0
                _ = v11
            default:
                if ((!(1e10 < 0.0)) && (42 > 127)) {
                    var v12 uint8 = 1
                    switch "\n\t" {
                    case "foo bar":
                        v5 = append(v5, fi6)
                    case "a":
                        v7 = (^(v7))
                    }
                    _ = v12
                }
            }
            _ = fi6
            _ = v7
        }
        _ = len(v5)
        _ = fi4
        _ = v5
    }
    switch 4294967295 {
    case 32768:
        switch 9223372036854775807 {
        case -114:
            var v13 uint8 = 2
            _ = v13
        case 32768:
            if (false && false) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 float32 = (1.5 * 0.0)
                    _ = fi14
                    _ = v15
                }
            }
        }
    case 4294967295:
        v16 := GenS2{F0: 42, F1: 100, F2: 18446744073709551615}
        _ = v16
    default:
        var v17 int16 = 32767
        _ = v17
    }
    return 127
}

func main() {
    var v1 bool = (2 >= 1)
    v2 := make(map[int]bool)
    var v3 int32 = 916
    switch ("a" + "hello") {
    case "":
        v1 = v1
    }
    v1 = false
    var v18 string = ""
    if false {
        _ = len(v2)
    }
    var v19 uint32 = 2147483647
    fmt.Println(v19, v19)
    if (42 < 708) {
        v20 := []int{-265, 0}
        _ = v20
    } else {
        var v21 uint = 892
        v22 := []int{0, -9223372036854775808, 42, int(v21)}
        delete(v2, 0)
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
}
