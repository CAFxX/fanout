package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            v1 = (0 <= 2147483648)
            var v7 bool = ("a" <= "\n\t")
            fmt.Println(fi5)
            _ = fi6
            _ = v7
        }
        v1 = (127 > 1)
        var v8 bool = (true && v1)
        _ = fi5
        _ = v8
    }
    switch -9223372036854775808 {
    case -1:
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 211}
            v11 := GenS1{F0: 1}
            if true {
                if false {
                    fmt.Println(fi9, v1)
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v13 := GenS1{F0: 0}
                        fmt.Println(fi9, fi9)
                        var v14 int = (fi9 + fi12)
                        var v15 uint64 = 1
                        v15 = (v15 >> ((4294967295) & 63))
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                } else {
                    var v16 float64 = 1e10
                    var v17 int64 = -9223372036854775808
                    v17 = (v17 >> ((255) & 63))
                    _ = v16
                    _ = v17
                }
                var v18 bool = true
                v1 = false
                _ = v18
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    case 9223372036854775807:
        _ = len(v3)
    case -355:
        fmt.Println(v1)
    }
    var v19 uint = 419
    fmt.Println(v19)
    v1 = (v1 && ("" >= "hello"))
    var v20 bool = (!(!v1))
    for fi21 := 0; fi21 < 16; fi21++ {
        delete(v4, 0)
        var v22 uint = v19
        switch 128 {
        case 2147483647:
            v23 := GenS1{F0: 1}
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    v24 := GenS2{F0: "", F1: 204}
    v1 = (v20 || (-32768 > 1000))
    v20 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
