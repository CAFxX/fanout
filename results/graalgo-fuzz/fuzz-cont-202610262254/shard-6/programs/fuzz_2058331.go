package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 uint; F2 int8 }
func genFn0() int16 {
    var v5 uint8 = 112
    for fi6 := 0; fi6 < 16; fi6++ {
        v5 -= v5
        if (-5 <= -207) {
            v5 -= (v5 * (^(v5)))
        }
        _ = fi6
    }
    var v7 uint16 = 65535
    v8 := GenS1{F0: 1000, F1: "\n\t"}
    v9 := GenS2{F0: (-32768 > 762), F1: 4294967295}
    _ = v5
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.0
    v3 := []int{-9223372036854775808, 9223372036854775807, 0, 0}
    v4 := make(map[uint32]bool)
    var v10 int = 233
    v10 = v10
    var v11 int32 = 852
    for fi12 := 0; fi12 < 12; fi12++ {
        if true {
            v13 := GenS2{F0: true, F1: 1}
            var v14 float32 = 1.5
            _ = len(v3)
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: false, F1: 18446744073709551615}
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        v10 -= fi16
        v11 = -2147483648
        _ = fi16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v2 += v2
        v18 := GenS3{F0: 1, F1: 929}
        var v19 int32 = 1
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 string = ("a" + (("a" + "a") + ""))
    v2 = v2
    fmt.Println(v20)
    if (false || (!(112 != 1))) {
        for fi21 := 0; fi21 < 7; fi21++ {
            v10 = fi21
            for fi22 := 0; fi22 < 8; fi22++ {
                v11 = 0
                _ = fi22
            }
            v20 = v20
            _ = fi21
        }
    } else {
        var v23 uint = 0
        _ = v23
    }
    v2 -= v2
    v10 = v10
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 int64 = -458
        for fi26 := 0; fi26 < 16; fi26++ {
            fmt.Println(fi24, fi26)
            var v27 uint16 = 780
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 14; fi28++ {
            if (v20 >= (((v20 + "foo bar") + "foo bar") + v20)) {
                if ((v20 + v20) > "hello") {
                    var v29 int = v10
                    var v30 uint64 = 4294967295
                    _ = v29
                    _ = v30
                } else {
                    fmt.Println(v2, fi28, fi24)
                }
                v2 = v2
                var v31 uint64 = 42
                _ = v31
            } else {
                v32 := GenS1{F0: genFn0(), F1: (v20 + v20)}
                if true {
                    var v33 bool = (v20 != v20)
                    _ = v33
                }
                v34 := GenS4{F0: fi24}
                var v35 float64 = (v2 - v2)
                _ = v32
                _ = v34
                _ = v35
            }
            var v36 float64 = 1.5
            _ = fi28
            _ = v36
        }
        _ = fi24
        _ = v25
    }
    if true {
        var v37 int16 = 1000
        v38 := GenS5{F0: 1, F1: 177, F2: 88}
        if (!false) {
            switch (v10 << ((0) & 63)) {
            case -9223372036854775808:
                for fi39 := 0; fi39 < 15; fi39++ {
                    delete(v4, 2147483647)
                    _ = fi39
                }
            case 32768:
                var v40 int8 = 100
                _ = v40
            }
        }
        _ = v37
        _ = v38
    }
    v10 += 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
