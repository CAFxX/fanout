package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16 }

func main() {
    var v1 string = "hello"
    var v2 float64 = float64((float32(0.1) - 1.5))
    v3 := make(map[string]int)
    var v4 uint16 = 42
    var v5 int16 = 606
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: -128}
    delete(v3, "a")
    if false {
        v2 += v2
        fmt.Println(v1)
    }
    v8 := GenS1{F0: -1}
    v2 = -2.25
    var v9 uint64 = 42
    switch -102 {
    case 4294967295:
        v10 := []int{}
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := []int{(-(fi11)), (fi12 | (-(fi11)))}
            _ = fi12
            _ = v13
        }
        switch (((v1 + "") + v1) + v1) {
        case "":
            for fi14 := 0; fi14 < 1; fi14++ {
                v4 = (^(v4))
                v5 += (v5 ^ v5)
                switch fi11 {
                case 11:
                    fmt.Println(v2, v5, v4)
                case 2147483647:
                    if false {
                        var v15 int = -2147483648
                        _ = v15
                    }
                case 814:
                    v2 -= 1.5
                }
                _ = fi14
            }
        }
        fmt.Println(v2, v1, v9)
        _ = fi11
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int16 = 1000
        for fi18 := 0; fi18 < 17; fi18++ {
            switch "foo bar" {
            case "foo bar":
                v19 := []int{int(v4), fi16, fi18, 0}
                _ = v19
            case "\n\t":
                v20 := v3[(("foo bar" + ((v1 + "hello") + "\n\t")) + ((v1 + v1) + v1))]
                _ = v20
            case "":
                var v21 int8 = -35
                _ = v21
            default:
                fmt.Println(v9, v9)
            }
            switch fi18 {
            case 582:
                v5 = v17
            }
            _ = fi18
        }
        v1 = v1
        _ = fi16
        _ = v17
    }
    v22 := GenS2{F0: (1 != 42)}
    v9 = v9
    fmt.Println(v9, v9)
    switch "" {
    case "\n\t":
        if false {
            for fi23 := 0; fi23 < 12; fi23++ {
                if (v1 < (("foo bar" + v1) + "a")) {
                    var v24 uint16 = v4
                    _ = v24
                }
                v25 := []int{fi23, (fi23 + 32768), fi23, (fi23 - (fi23 & fi23))}
                _ = fi23
                _ = v25
            }
            for fi26 := 0; fi26 < 2; fi26++ {
                fmt.Println(v9)
                _ = fi26
            }
            if false {
                fmt.Println(v4, v2, v1)
            }
        }
    }
    v5 -= v5
    v27 := GenS3{F0: 0, F1: (^(v5))}
    var v28 uint32 = 667
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
