package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 18446744073709551615
    v3 := make(map[string]string)
    var v4 int64 = -9223372036854775808
    var v5 float64 = -2.25
    var v6 bool = false
    if (("a" < "\n\t") || v6) {
        v7 := v1[(("a" + "hello") + "\n\t")]
        v7 -= v7
        v8 := []int{}
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 226}
        _ = len(v3)
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint = v2
        if v6 {
            v6 = v6
        }
        _ = fi10
        _ = v11
    }
    v2 += (v2 << (((v2 + v2)) & 63))
    v12 := GenS1{F0: 127}
    var v13 uint32 = 4294967295
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            if false {
                v16 := GenS2{F0: 0}
                _ = len(v3)
                _ = v16
            } else {
                v17 := GenS3{F0: 187}
                v13 = v13
                var v18 bool = (v6 || (!true))
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        var v19 int16 = 32767
        _ = fi14
        _ = v19
    }
    v20 := GenS4{F0: v6}
    var v21 uint8 = 1
    v5 += v5
    v6 = (!(v6 || (!v6)))
    for fi22 := 0; fi22 < 20; fi22++ {
        switch (fi22 >> ((v2) & 63)) {
        case -2147483648:
            v21 = (v21 ^ 128)
        case -418:
            v23 := GenS2{F0: 18446744073709551615}
            _ = v23
        case 591:
            for fi24 := 0; fi24 < 8; fi24++ {
                v4 += (v4 + v4)
                v25 := GenS2{F0: uint64(v21)}
                if v6 {
                    _ = len(v1)
                    v21 = (^(v21))
                    var v26 uint8 = 127
                    _ = len(v3)
                    _ = v26
                } else {
                    switch "" {
                    case "foo bar":
                        v27 := v3["a"]
                        _ = v27
                    }
                    var v28 uint32 = (v13 ^ v13)
                    var v29 int64 = (v4 ^ int64(v2))
                    var v30 uint32 = v13
                    fmt.Println(v30)
                    v31 := GenS5{F0: v28}
                    var v32 float64 = v5
                    v33 := GenS5{F0: (^(v13))}
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                var v34 uint32 = v13
                v35 := GenS6{F0: (^(v2)), F1: -1}
                _ = fi24
                _ = v25
                _ = v34
                _ = v35
            }
        }
        v6 = v6
        _ = fi22
    }
    if true {
        var v36 float64 = v5
        v37 := GenS5{F0: (^(v13))}
        _ = v36
        _ = v37
    }
    var v38 uint8 = v21
    var v39 int32 = 42
    v40 := []int{9223372036854775807, 128, 4294967295, 128}
    v39 -= int32(v38)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
