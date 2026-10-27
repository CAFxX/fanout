package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint16; F2 int }
func genFn0(p0 int32, p1 uint32, p2 uint64) int64 {
    p1 = p1
    switch "a" {
    case "":
        for fi4 := 0; fi4 < 9; fi4++ {
            v5 := []int{fi4, (^(fi4))}
            for fi6 := 0; fi6 < 16; fi6++ {
                fmt.Println(p2, p1, fi4)
                var v7 float64 = 0.0
                if (128 >= 0) {
                    if len(v5) > 0 {
                        v5[0] = (fi4 << ((0) & 63))
                    }
                } else {
                    var v8 float64 = -2.25
                    _ = v8
                }
                var v9 int64 = 9223372036854775807
                _ = fi6
                _ = v7
                _ = v9
            }
            _ = fi4
            _ = v5
        }
    case "foo bar":
        v10 := GenS1{F0: 128}
        _ = v10
    }
    var v11 int32 = (p0 & (p0 >> ((18446744073709551615) & 63)))
    var v12 uint8 = 58
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 1
}
func genFn1(p0 uint64, p1 bool, p2 uint32) uint8 {
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint64 = 42
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := GenS1{F0: uint8(p0)}
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 uint8 = 2
    var v2 uint16 = 1
    var v3 bool = false
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint = 484
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 uint32 = uint32(v18)
            _ = fi19
            _ = v20
        }
        v1 = v1
        _ = fi17
        _ = v18
    }
    if (!true) {
        v21 := GenS2{F0: 2147483647}
        v22 := GenS2{F0: 731}
        _ = v21
        _ = v22
    }
    var v23 bool = v3
    if true {
        v24 := GenS2{F0: -573}
        _ = v24
    }
    var v25 uint = 0
    v26 := []int{994, 127, -535}
    v26 = append(v26, 127)
    for fi27 := 0; fi27 < 6; fi27++ {
        for fi28 := 0; fi28 < 10; fi28++ {
            if (!v23) {
                if len(v26) > 0 {
                    v26[0] = 128
                }
                switch (("\n\t" + "a") + ("\n\t" + "\n\t")) {
                case "foo bar":
                    v29 := GenS3{F0: fi28}
                    _ = v29
                case "\n\t":
                    var v30 uint8 = (v1 / (((v1 % ((216) | 1))) | 1))
                    _ = v30
                }
                v1 = 2
                switch "a" {
                case "hello":
                    var v31 uint32 = 940
                    v2 = 323
                    _ = v31
                case "a":
                    _ = len(v26)
                }
                fmt.Println(v2, v23)
            }
            if (1 != 0) {
                if len(v26) > 0 {
                    v26[0] = (fi28 >> ((v25) & 63))
                }
                var v32 int64 = 9223372036854775807
                var v33 uint = (^(v25))
                _ = v32
                _ = v33
            }
            _ = fi28
        }
        v34 := GenS4{F0: 1, F1: v2, F2: fi27}
        _ = fi27
        _ = v34
    }
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    var v35 uint64 = 416
    _ = len(v26)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v35
}
