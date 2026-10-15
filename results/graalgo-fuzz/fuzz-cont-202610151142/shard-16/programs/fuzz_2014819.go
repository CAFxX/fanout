package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0(p0 uint8) string {
    v6 := GenS1{F0: 65535}
    p0 = 42
    fmt.Println(p0)
    _ = p0
    _ = v6
    return ("foo bar" + "foo bar")
}
func genFn1(p0 int16, p1 int8, p2 int16) int {
    p1 = p1
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 uint32 = 65535
    v4 := make(map[uint32]int)
    var v5 int = 4294967295
    delete(v2, ((("a" + "") + (("foo bar" + ("a" + "foo bar")) + (("\n\t" + "") + "\n\t"))) + ("a" + (("" + "a") + "hello"))))
    var v7 uint16 = 0
    var v8 float32 = (-2.25 / float32((3.14159 / 3.14159)))
    var v9 uint64 = 0
    var v10 uint16 = (v7 * v7)
    v3 = v3
    var v11 float32 = -2.25
    fmt.Println(v7)
    fmt.Println(v3, v3)
    var v12 int32 = 2147483647
    v13 := []int{-2147483648}
    v12 += v12
    switch (v5 / (((v5 - (v5 >> ((1) & 63)))) | 1)) {
    case -9223372036854775808:
        delete(v4, 519)
    case 32767:
        for fi14 := 0; fi14 < 4; fi14++ {
            switch genFn1(-1, 1, 1) {
            case 962:
                var v15 uint32 = (v3 * v3)
                _ = v15
            case 128:
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 int64 = 725
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi14
        }
    case 443:
        for fi18 := 0; fi18 < 10; fi18++ {
            for fi19 := 0; fi19 < 5; fi19++ {
                v5 = genFn1(979, 100, 1000)
                _ = fi19
            }
            v9 = (v9 & (v9 / ((v9) | 1)))
            if true {
                var v20 uint64 = v9
                v21 := GenS2{F0: true}
                _ = v20
                _ = v21
            }
            _ = fi18
        }
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int32 = v12
        for fi24 := 0; fi24 < 17; fi24++ {
            v25 := GenS3{F0: 758, F1: (v23 / (((v23 ^ 42)) | 1))}
            v3 = (v3 << ((42) & 63))
            fmt.Println(v12)
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    v10 = (v7 - v7)
    if (190 < -1) {
        switch (("a" + ("hello" + ("a" + "foo bar"))) + "foo bar") {
        case "hello":
            delete(v4, v3)
        case "":
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 int8 = 127
                _ = fi26
                _ = v27
            }
        default:
            fmt.Println(v5, v11, v12)
        }
    }
    fmt.Println(v10, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
