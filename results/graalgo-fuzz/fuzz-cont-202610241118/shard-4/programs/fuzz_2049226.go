package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 65535
    var v4 int32 = v1
    v5 := GenS1{F0: 2, F1: 0}
    var v6 uint8 = 172
    v7 := GenS2{F0: 128}
    fmt.Println(v6)
    var v8 int8 = 0
    var v9 string = "foo bar"
    var v10 uint64 = 165
    for fi11 := 0; fi11 < 11; fi11++ {
        if true {
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 float32 = (0.1 / 1.5)
                _ = fi12
                _ = v13
            }
            v14 := []int{701, fi11, 32768}
            _ = v14
        }
        _ = fi11
    }
    var v15 bool = true
    for fi16 := 0; fi16 < 5; fi16++ {
        fmt.Println(v4, v10, v1)
        _ = fi16
    }
    fmt.Println(v10)
    fmt.Println(v9)
    for fi17 := 0; fi17 < 7; fi17++ {
        fmt.Println(v10)
        _ = fi17
    }
    switch 2147483647 {
    case 9223372036854775807:
        v3 = (v3 - (^(v3)))
    }
    v10 += (v10 + v10)
    switch ((v9 + ((v9 + v9) + "\n\t")) + (("hello" + (((((v9 + "a") + v9) + (v9 + "\n\t")) + (("" + v9) + ((("foo bar" + "foo bar") + "") + "a"))) + v9)) + (("\n\t" + v9) + (v9 + v9)))) {
    case "a":
        v18 := GenS2{F0: 1}
        _ = v18
    }
    v19 := GenS3{F0: "foo bar", F1: 510}
    var v20 bool = (-1 >= 1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
