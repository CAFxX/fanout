package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int8; F1 int8; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint }
func genFn0(p0 int) uint32 {
    var v4 int32 = 1
    for fi5 := 0; fi5 < 19; fi5++ {
        v4 = (v4 % (((v4 - v4)) | 1))
        p0 = -9223372036854775808
        _ = fi5
    }
    var v6 int = (p0 % (((-(p0))) | 1))
    v7 := []int{(p0 / ((p0) | 1)), (p0 * p0)}
    _ = p0
    _ = v4
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1() int8 {
    var v8 int32 = 42
    v8 = -2147483648
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]string)
    var v3 string = "\n\t"
    v1 -= v1
    v9 := GenS1{F0: -32768, F1: 0}
    var v10 uint32 = 2
    switch 4294967295 {
    case 694:
        for fi11 := 0; fi11 < 10; fi11++ {
            v1 -= (v1 & v1)
            _ = fi11
        }
    case 127:
        v1 = v1
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := []int{fi13, fi12}
            if true {
                fmt.Println(v10, v3, v1)
            }
            v15 := GenS1{F0: 1, F1: 761}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 string = v3
        if (!true) {
            v3 = ("" + "hello")
        } else {
            v18 := GenS2{F0: 1, F1: fi16}
            switch "\n\t" {
            case "a":
                delete(v2, (fi16 + (fi16 + 32767)))
            case "hello":
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 int = fi16
                    delete(v2, -2147483648)
                    _ = fi19
                    _ = v20
                }
            }
            v21 := []int{-9223372036854775808, (fi16 * fi16)}
            _ = v18
            _ = v21
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int8 = -128
            v24 := GenS3{F0: (v23 / ((1) | 1)), F1: (v23 ^ 9), F2: (^(v1))}
            var v25 uint = 42
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi16
        _ = v17
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v3 = v3
        _ = fi26
    }
    v27 := GenS4{F0: 1, F1: 661}
    v1 -= v1
    if (!true) {
        switch 1 {
        case 4294967295:
            fmt.Println(v3, v1)
        }
        var v28 int32 = 948
        if true {
            for fi29 := 0; fi29 < 10; fi29++ {
                delete(v2, -921)
                _ = fi29
            }
            _ = len(v2)
        } else {
            for fi30 := 0; fi30 < 17; fi30++ {
                v31 := GenS3{F0: -1, F1: -13, F2: 1}
                _ = fi30
                _ = v31
            }
        }
        _ = v28
    } else {
        v32 := GenS4{F0: 0, F1: 252}
        v3 = ("a" + ((v3 + "") + (v3 + v3)))
        _ = v32
    }
    switch v3 {
    case "foo bar":
        var v33 uint = 0
        _ = v33
    }
    var v34 int = 32768
    if false {
        fmt.Println(v1, v10)
        var v35 int64 = 9223372036854775807
        _ = v35
    } else {
        if (true && (v3 < v3)) {
            v36 := []int{}
            _ = v36
        }
        fmt.Println(v1, v1, v34)
        var v37 int32 = -1
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v9
}
