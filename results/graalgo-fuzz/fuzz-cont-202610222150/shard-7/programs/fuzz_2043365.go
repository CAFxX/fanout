package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int16; F1 uint64; F2 bool }
type GenS4 struct { F0 bool }

func main() {
    var v1 string = "foo bar"
    var v2 bool = true
    v3 := make(map[int]uint32)
    fmt.Println(v2, v1)
    v4 := GenS1{F0: -2147483648, F1: (v1 + "hello")}
    fmt.Println(v2)
    if true {
        v2 = v2
        var v5 uint32 = 4294967295
        if false {
            v6 := GenS2{F0: 2, F1: 128}
            v7 := GenS1{F0: 42, F1: ("a" + v1)}
            var v8 bool = v2
            _ = v6
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        var v9 uint64 = 897
        v10 := GenS3{F0: 1000, F1: 741, F2: v2}
        for fi11 := 0; fi11 < 12; fi11++ {
            if true {
                v1 = ((("a" + v1) + v1) + v1)
            } else {
                if v2 {
                    v12 := GenS3{F0: 370, F1: (v9 / ((v9) | 1)), F2: (!(!v2))}
                    delete(v3, int(v9))
                    _ = v12
                } else {
                    if (42 >= -1) {
                        v13 := GenS4{F0: (!v2)}
                        v14 := GenS4{F0: (913 < 1)}
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint = 299
                        var v16 string = (v1 + (v1 + "a"))
                        v17 := []int{(fi11 % ((1) | 1)), (fi11 - fi11), fi11}
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 5; fi18++ {
                        var v19 uint8 = 127
                        delete(v3, fi11)
                        _ = fi18
                        _ = v19
                    }
                    var v20 int = fi11
                    _ = v20
                }
                if (!(!v2)) {
                    fmt.Println(v9)
                    fmt.Println(v1)
                    switch fi11 {
                    case 42:
                        if (true && v2) {
                            v21 := GenS1{F0: 881, F1: "\n\t"}
                            fmt.Println(fi11, v9)
                            _ = v21
                        } else {
                            v9 = v9
                        }
                    case 32767:
                        v22 := []int{fi11}
                        _ = v22
                    }
                }
                fmt.Println(v1, fi11)
            }
            v23 := GenS2{F0: 32767, F1: -9223372036854775808}
            var v24 int8 = 100
            _ = fi11
            _ = v23
            _ = v24
        }
        _ = v10
        _ = v9
    }
    v1 = (v1 + v1)
    switch ("" + v1) {
    case "":
        var v25 int32 = -538
        _ = v25
    default:
        var v26 int8 = 100
        _ = v26
    }
    v27 := v3[2147483647]
    v3[32768] = (v27 / (((v27 >> ((0) & 63))) | 1))
    var v28 int16 = 459
    v29 := GenS2{F0: 42, F1: 128}
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 string = (v1 + "hello")
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
