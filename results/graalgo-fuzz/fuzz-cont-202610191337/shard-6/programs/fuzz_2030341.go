package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int32; F1 int8; F2 bool }
type GenS4 struct { F0 uint64; F1 int8 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: 2147483648, F1: 42, F2: 9223372036854775807}
            var v7 int64 = 9223372036854775807
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: 65535, F1: 2, F2: -9223372036854775808}
        fmt.Println(fi4)
        _ = fi4
        _ = v8
    }
    switch ((((("a" + "hello") + ((("\n\t" + (("foo bar" + ("hello" + "")) + "a")) + "a") + ("a" + ("" + "")))) + "foo bar") + ("a" + "hello")) + (("\n\t" + (("" + ("hello" + "\n\t")) + "a")) + ("hello" + ""))) {
    case "foo bar":
        for fi9 := 0; fi9 < 2; fi9++ {
            switch (("foo bar" + "") + "foo bar") {
            case "a":
                var v10 uint = 1
                _ = v10
            case "\n\t":
                v11 := GenS2{F0: 32767, F1: -1}
                _ = v11
            }
            fmt.Println(fi9)
            var v12 uint64 = 42
            _ = fi9
            _ = v12
        }
    }
    var v13 int8 = 100
    v14 := GenS2{F0: 420, F1: -1}
    _ = v13
    _ = v14
    return 1
}
func genFn1(p0 int64) bool {
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS2{F0: 32767, F1: 42}
        for fi17 := 0; fi17 < 8; fi17++ {
            p0 = 729
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := []int{fi18, (fi18 | -9223372036854775808), (fi18 ^ fi18)}
        v20 := GenS2{F0: 2, F1: 1000}
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS2{F0: 621, F1: int16(p0)}
            _ = fi21
            _ = v22
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    return true
}

func main() {
    var v1 uint = 42
    var v2 int64 = 190
    var v3 int64 = 42
    v23 := GenS2{F0: 555, F1: genFn0()}
    if (v2 >= (v2 << ((v1) & 63))) {
        var v24 int32 = 1
        v25 := GenS1{F0: 65535, F1: 2, F2: v3}
        _ = v24
        _ = v25
    } else {
        var v26 int32 = 1
        var v27 int = 128
        _ = v26
        _ = v27
    }
    var v28 uint8 = 42
    var v29 bool = false
    var v30 uint8 = (v28 ^ v28)
    v31 := GenS3{F0: 0, F1: 0, F2: v29}
    v32 := []int{-1, -724, 321}
    v32 = append(v32, 54)
    var v33 string = ""
    v34 := GenS2{F0: 542, F1: 326}
    var v35 uint = (v1 >> (((v1 + (v1 + (v1 + (v1 % (((^(v1))) | 1)))))) & 63))
    var v36 uint = (^(v1))
    for fi37 := 0; fi37 < 15; fi37++ {
        switch fi37 {
        case -774:
            v38 := GenS4{F0: 18446744073709551615, F1: 0}
            _ = v38
        }
        _ = fi37
    }
    v39 := GenS2{F0: uint16(v36), F1: 0}
    var v40 uint16 = 336
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v40
}
