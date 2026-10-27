package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 string }
func genFn0() string {
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = 3.14159
        _ = fi7
        _ = v8
    }
    v6 = (v6 ^ (v6 * -9223372036854775808))
    _ = v6
    return "foo bar"
}
func genFn1() uint {
    var v9 bool = ("hello" >= "hello")
    fmt.Println(v9)
    v9 = (1 >= 2147483647)
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 float64 = 1e10
        fmt.Println(fi10, v11)
        v12 := []int{(fi10 / ((fi10) | 1))}
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v9
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 937
    v3 := make(map[uint32]int)
    var v4 int32 = 42
    v5 := []uint32{0, 1, 0, 4294967295}
    v4 += (v4 / (((-(v4))) | 1))
    v13 := v3[141]
    var v14 uint32 = 65535
    if true {
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS1{F0: "\n\t"}
            _ = fi15
            _ = v16
        }
        if (3.14159 <= -2.25) {
            for fi17 := 0; fi17 < 9; fi17++ {
                v13 -= fi17
                var v18 float64 = float64(3.14159)
                fmt.Println(v2, v14, v18)
                _ = fi17
                _ = v18
            }
        }
        v19 := v3[(^(v14))]
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 string = "hello"
        if len(v5) > 0 {
            v5[0] = 449
        }
        _ = fi20
        _ = v21
    }
    switch (v13 - (^(v13))) {
    case 1:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 32767:
        fmt.Println(v14, v14, v13)
    }
    var v22 uint32 = (v14 + 209)
    if false {
        v23 := GenS2{F0: (v13 / ((v13) | 1)), F1: 743, F2: v22}
        var v24 int32 = int32(v13)
        fmt.Println(v24, v22)
        _ = v23
        _ = v24
    } else {
        v25 := GenS3{F0: v13}
        switch "hello" {
        case "":
            var v26 float64 = -2.25
            _ = v26
        case "\n\t":
            v27 := GenS2{F0: 842, F1: 255, F2: v14}
            _ = v27
        }
        _ = v25
    }
    var v28 uint64 = 766
    v22 = (v14 - v22)
    v29 := GenS1{F0: (("hello" + (("hello" + "") + "")) + ((("foo bar" + "") + ("\n\t" + (("foo bar" + "hello") + ("foo bar" + "foo bar")))) + "a"))}
    fmt.Println(v4, v28, v4)
    _ = len(v3)
    var v30 uint8 = 255
    v31 := GenS4{F0: (v28 | v28), F1: ""}
    switch (-(v13)) {
    case -1:
        var v32 int64 = v2
        _ = v32
    case 128:
        var v33 uint32 = 640
        _ = v33
    case 32767:
        v13 += (v13 >> ((genFn1()) & 63))
    }
    var v34 uint64 = (^(v28))
    fmt.Println(v22, v28, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
