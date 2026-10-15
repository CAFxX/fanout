package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int16 }
func genFn0() string {
    if true {
        if (23 > 1) {
            switch ("hello" + "\n\t") {
            case "a":
                v6 := []int{2147483647, 9223372036854775807, 0, -9223372036854775808}
                _ = v6
            case "\n\t":
                var v7 int32 = 154
                _ = v7
            case "foo bar":
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 int64 = -969
                    if true {
                        v10 := GenS1{F0: -2147483648}
                        _ = v10
                    } else {
                        var v11 int64 = -9223372036854775808
                        if false {
                            v12 := GenS1{F0: -2147483648}
                            _ = v12
                        }
                        var v13 uint8 = 1
                        _ = v11
                        _ = v13
                    }
                    _ = fi8
                    _ = v9
                }
            }
            v14 := GenS1{F0: -578}
            var v15 uint32 = 4294967295
            _ = v14
            _ = v15
        }
        var v16 uint32 = 42
        _ = v16
    } else {
        var v17 uint32 = 904
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            if false {
                var v20 uint64 = 2
                var v21 float64 = 0.0
                v20 -= v20
                v20 -= (v20 - 0)
                _ = v20
                _ = v21
            }
            _ = fi19
        }
        _ = fi18
    }
    var v22 uint32 = 42
    var v23 uint64 = 4294967295
    v22 += v22
    _ = v22
    _ = v23
    return "hello"
}
func genFn1(p0 int32, p1 uint16) uint32 {
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS1{F0: (p0 * -2147483648)}
        var v26 float32 = 1.5
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := []int{fi27, (^(fi24)), 536}
            var v29 int = (-(fi24))
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    var v30 uint = 0
    fmt.Println(v30)
    v31 := GenS2{F0: 961, F1: 65535, F2: 802}
    v32 := []int{-692, 32767}
    _ = p0
    _ = p1
    _ = v30
    _ = v31
    _ = v32
    return 65535
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 2
    var v3 string = ("a" + ("" + "hello"))
    var v4 bool = false
    v5 := []uint32{1}
    fmt.Println(v2, v1, v1)
    var v33 float64 = 0.1
    var v34 string = (v3 + ("a" + "a"))
    v1 = v1
    for fi35 := 0; fi35 < 20; fi35++ {
        v3 = "a"
        var v36 uint = 4294967295
        v2 = (v2 & uint8(v1))
        _ = fi35
        _ = v36
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v1, v4)
    fmt.Println(v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
