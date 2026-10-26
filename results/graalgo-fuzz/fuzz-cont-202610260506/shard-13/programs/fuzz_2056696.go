package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    var v4 int32 = 0
    var v5 string = ("hello" + "")
    v6 := make(map[string]string)
    if (-32768 < 42) {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 32767, F1: 236}
            _ = fi7
            _ = v8
        }
        v5 = ("foo bar" + (("a" + v5) + ""))
    } else {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 bool = true
            v10 = (1000 != 0)
            if (v10 || v10) {
                var v11 int32 = (v4 + 1)
                v12 := GenS1{F0: 0, F1: 814}
                var v13 int32 = (v11 ^ 42)
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
        }
    }
    v14 := []int{-997, -9223372036854775808, 32768, 128}
    if false {
        v4 -= v4
        _ = len(v1)
        v15 := GenS1{F0: 42, F1: 42}
        _ = v15
    } else {
        v5 = ((((v5 + (v5 + "a")) + v5) + (("foo bar" + "foo bar") + "a")) + "")
        switch (v5 + ("\n\t" + ((("\n\t" + v5) + "hello") + "hello"))) {
        case "":
            _ = len(v3)
        default:
            v16 := GenS1{F0: 398, F1: 2147483647}
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v4 += v4
        v5 = v5
        _ = fi17
    }
    var v18 uint64 = 0
    for fi19 := 0; fi19 < 1; fi19++ {
        v5 = v5
        if (true && true) {
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 float64 = 1e10
                if len(v14) > 0 {
                    v14[0] = fi19
                }
                v22 := GenS1{F0: 0, F1: 2}
                _ = fi20
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 14; fi23++ {
                switch v5 {
                case "\n\t":
                    delete(v2, (-(fi19)))
                case "":
                    fmt.Println(v4)
                default:
                    v5 = v5
                }
                v5 = (v5 + "a")
                _ = fi23
            }
            for fi24 := 0; fi24 < 17; fi24++ {
                if (!(true || (false && (!true)))) {
                    delete(v6, (v5 + (("" + "\n\t") + (v5 + v5))))
                    v25 := GenS2{F0: 67, F1: 2147483648}
                    v26 := GenS3{F0: 387, F1: 127, F2: 65535}
                    _ = v25
                    _ = v26
                }
                _ = fi24
            }
        }
        _ = fi19
    }
    var v27 int8 = -1
    fmt.Println(v4, v5, v5)
    delete(v3, 2147483647)
    fmt.Println(v5, v18, v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
