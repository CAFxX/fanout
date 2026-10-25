package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 int16 }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }
func genFn0(p0 uint64) uint16 {
    for fi5 := 0; fi5 < 8; fi5++ {
        switch fi5 {
        case 2147483647:
            var v6 uint8 = 128
            _ = v6
        case 128:
            if ((!false) && true) {
                fmt.Println(p0)
            } else {
                v7 := GenS1{F0: 128, F1: 394}
                p0 += p0
                _ = v7
            }
        }
        _ = fi5
    }
    var v8 int8 = 101
    v9 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v8
    _ = v9
    return 702
}
func genFn1(p0 uint, p1 uint8, p2 int8) uint64 {
    fmt.Println(p2)
    v10 := GenS2{F0: p1, F1: "", F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 0
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]int)
    var v3 int = -284
    v4 := make(map[uint32]string)
    v3 = (v3 | -802)
    var v11 int = int(v1)
    v11 = (v11 & (-(v11)))
    if ((!false) && false) {
        if true {
            v12 := []int{}
            v11 += (v3 + (v3 % ((2147483647) | 1)))
            switch ("hello" + "hello") {
            case "\n\t":
                v11 = (v11 * v11)
            case "foo bar":
                fmt.Println(v3)
            case "hello":
                var v13 float32 = (((1e10 - (0.0 / ((0.1 * 1.5) - ((3.14159 / 1.5) - (0.1 / 0.1))))) / (0.1 - (((0.1 - 3.14159) / 0.1) + float32(1.5)))) + 0.0)
                _ = v13
            default:
                _ = len(v4)
            }
            _ = v12
        }
        v11 = (v11 % ((v3) | 1))
        fmt.Println(v11)
    }
    var v14 uint32 = 42
    var v15 int = 32768
    v16 := []int{}
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint64 = 2
        v19 := GenS3{F0: 487, F1: -962, F2: (v18 ^ (v18 ^ v18))}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: 2, F1: (v14 ^ v14)}
    if (!false) {
        switch v3 {
        case 637:
            switch ("\n\t" + "hello") {
            case "\n\t":
                fmt.Println(v1, v1, v15)
            }
        case 2147483647:
            v21 := GenS2{F0: 127, F1: "", F2: -1}
            _ = v21
        }
    }
    v14 = 983
    v3 += 4294967295
    var v22 uint16 = genFn0(4294967295)
    switch (^(v11)) {
    case -912:
        v23 := GenS2{F0: 253, F1: "", F2: 42}
        _ = v23
    case 1:
        switch "foo bar" {
        case "a":
            v24 := GenS1{F0: 1, F1: 1}
            _ = v24
        case "\n\t":
            var v25 uint16 = genFn0(1)
            _ = v25
        case "":
            v26 := GenS3{F0: 42, F1: -520, F2: 18446744073709551615}
            _ = v26
        default:
            for fi27 := 0; fi27 < 5; fi27++ {
                fmt.Println(v15)
                var v28 uint16 = v22
                _ = fi27
                _ = v28
            }
        }
    default:
        for fi29 := 0; fi29 < 16; fi29++ {
            v30 := GenS2{F0: 255, F1: ("\n\t" + (((("\n\t" + "hello") + "\n\t") + "\n\t") + "foo bar")), F2: 1}
            _ = fi29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 int32 = 1
        switch ((("hello" + "hello") + "hello") + "hello") {
        case "":
            var v33 uint = 757
            _ = v33
        }
        delete(v2, ((("hello" + "\n\t") + "\n\t") + ("foo bar" + ("a" + "foo bar"))))
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
}
