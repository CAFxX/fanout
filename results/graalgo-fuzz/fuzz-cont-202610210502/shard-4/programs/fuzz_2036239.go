package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 923
    var v3 string = ("\n\t" + "hello")
    var v4 uint16 = 32767
    var v5 uint = 0
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            if true {
                fmt.Println(v2)
                v8 := GenS1{F0: 993, F1: 217, F2: 127}
                _ = v8
            }
            switch ((v3 + (v3 + v3)) + v3) {
            case "a":
                if (!(-307 <= 366)) {
                    var v9 int32 = 615
                    v9 = v2
                    v10 := GenS1{F0: 267, F1: 154, F2: 2147483647}
                    _ = v10
                    _ = v9
                } else {
                    v11 := []int{fi7}
                    _ = v11
                }
            }
            v12 := GenS2{F0: 100}
            _ = fi7
            _ = v12
        }
        _ = fi6
    }
    v4 -= v4
    v13 := []int{-61, 206, 61, -195}
    for fi14 := 0; fi14 < 18; fi14++ {
        switch fi14 {
        case 0:
            var v15 int = 128
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 int = (fi14 | (fi16 % (((^(fi16))) | 1)))
            v18 := GenS3{F0: -1, F1: 1}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v4 = v4
        _ = fi19
    }
    v20 := GenS1{F0: 714, F1: 255, F2: 32767}
    var v21 float64 = -2.25
    v22 := GenS1{F0: 997, F1: 2, F2: 128}
    v5 = v5
    for fi23 := 0; fi23 < 5; fi23++ {
        v21 = 1.5
        _ = fi23
    }
    v3 = ("hello" + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
