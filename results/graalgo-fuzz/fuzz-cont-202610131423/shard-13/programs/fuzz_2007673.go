package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
func genFn0() uint {
    if false {
        v7 := GenS1{F0: true, F1: 255, F2: ("hello" + ("foo bar" + ""))}
        var v8 bool = (true || true)
        _ = v7
        _ = v8
    } else {
        var v9 uint64 = 412
        if true {
            v10 := []int{128, 32768, 42, 32768}
            switch "foo bar" {
            case "hello":
                v11 := GenS1{F0: (65535 < 255), F1: 2, F2: ("a" + (("foo bar" + ((("hello" + "hello") + "") + "")) + ("a" + "foo bar")))}
                _ = v11
            case "\n\t":
                v9 = (v9 ^ 42)
            case "":
                var v12 uint64 = 1
                _ = v12
            }
            _ = v10
        }
        if ((2147483648 == 814) && true) {
            v13 := GenS2{F0: 42}
            fmt.Println(v9)
            _ = v13
        }
        _ = v9
    }
    var v14 uint32 = 4294967295
    _ = v14
    return 42
}
func genFn1(p0 int32, p1 uint8) uint32 {
    p0 += (p0 << ((2) & 63))
    switch ("a" + "") {
    case "hello":
        v15 := GenS1{F0: (0 <= 4294967295), F1: 32767, F2: ""}
        _ = v15
    }
    p0 += p0
    var v16 float32 = -2.25
    _ = p0
    _ = p1
    _ = v16
    return 657
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 float32 = float32(1e10)
    var v4 float64 = 3.14159
    v5 := []int{4294967295, 42}
    v6 := make(map[uint32]string)
    var v17 uint64 = 2
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v4, v3, v4)
        v4 += -2.25
        _ = fi18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        v17 = (^(v17))
        _ = fi19
    }
    var v20 string = "foo bar"
    for fi21 := 0; fi21 < 2; fi21++ {
        v4 = 1e10
        var v22 bool = (false && true)
        v17 -= (^(v17))
        _ = fi21
        _ = v22
    }
    switch v20 {
    case "":
        var v23 bool = (!false)
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v17 = 2
        var v25 string = ("\n\t" + (v20 + ("a" + ((("foo bar" + v20) + "a") + ("\n\t" + "a")))))
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: 110, F1: 165, F2: 0}
    fmt.Println(v4, v4, v17)
    v5 = append(v5, -2147483648)
    var v27 int16 = 1
    v17 -= 2
    fmt.Println(v17)
    var v28 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
