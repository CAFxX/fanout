package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint8; F2 int16 }

func main() {
    var v1 uint64 = 42
    var v2 bool = false
    var v3 uint32 = 65535
    var v4 uint64 = (v1 << ((18446744073709551615) & 63))
    v5 := make(map[int]string)
    var v6 int16 = 1
    for fi7 := 0; fi7 < 10; fi7++ {
        v2 = (((v2 && v2) || v2) || v2)
        var v8 uint8 = 2
        _ = fi7
        _ = v8
    }
    fmt.Println(v3, v4, v2)
    switch 32768 {
    case 0:
        v2 = v2
    case -1:
        v9 := GenS1{F0: ("" + "hello")}
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS2{F0: 1}
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v13 := GenS3{F0: int16(v3)}
    var v14 uint16 = 255
    switch (("hello" + "\n\t") + "a") {
    case "hello":
        v15 := GenS4{F0: 930, F1: 1, F2: -472}
        _ = v15
    case "a":
        v16 := GenS2{F0: 127}
        _ = v16
    }
    v17 := GenS2{F0: uint8(v3)}
    fmt.Println(v14)
    v18 := v5[4294967295]
    if v2 {
        if (!(!v2)) {
            if v2 {
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 float64 = -2.25
                    _ = fi19
                    _ = v20
                }
                if v2 {
                    v21 := GenS1{F0: (v18 + v18)}
                    _ = v21
                } else {
                    v22 := v5[127]
                    v5[127] = v18
                    var v23 int64 = 42
                    _ = v22
                    _ = v23
                }
                fmt.Println(v14)
            } else {
                _ = len(v5)
                if false {
                    var v24 string = ("hello" + v18)
                    fmt.Println(v6)
                    _ = v24
                } else {
                    v2 = (42 != v6)
                    v2 = ((-2.25 < (float32(0.0) * 0.0)) && (!v2))
                }
                for fi25 := 0; fi25 < 2; fi25++ {
                    var v26 uint8 = 163
                    _ = fi25
                    _ = v26
                }
            }
        }
    }
    if true {
        v1 -= v4
        for fi27 := 0; fi27 < 18; fi27++ {
            v6 = 1
            _ = fi27
        }
    }
    fmt.Println(v4, v18)
    v28 := GenS2{F0: 42}
    v29 := []int{-2147483648}
    for fi30 := 0; fi30 < 10; fi30++ {
        v31 := GenS1{F0: (v18 + ("" + v18))}
        _ = fi30
        _ = v31
    }
    var v32 int32 = -280
    var v33 uint32 = (v3 ^ 175)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
