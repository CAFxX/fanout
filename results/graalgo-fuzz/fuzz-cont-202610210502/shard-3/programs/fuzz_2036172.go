package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 string; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 float32 = (3.14159 + -2.25)
    v2 := make(map[uint32]string)
    var v3 int16 = -32768
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint32 = 694
        v5 -= (^(v5))
        _ = fi4
        _ = v5
    }
    if false {
        v6 := []int{}
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 int = (fi7 % (((fi8 << ((2) & 63))) | 1))
                v10 := GenS1{F0: -1}
                var v11 float64 = 0.0
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi7
        }
        var v12 string = ("a" + "")
        _ = v12
        _ = v6
    } else {
        switch 9223372036854775807 {
        case 128:
            if (2 > 22) {
                var v13 float64 = 1e10
                switch (("a" + "") + "hello") {
                case "a":
                    fmt.Println(v1)
                case "hello":
                    var v14 uint32 = 65535
                    _ = v14
                case "foo bar":
                    var v15 float64 = v13
                    _ = v15
                }
                v3 = (v3 & v3)
                _ = v13
            } else {
                v1 = float32((0.0 * -2.25))
                v3 = (v3 + v3)
                v16 := GenS2{F0: 1, F1: "foo bar", F2: (((v1 - v1) / 1.5) <= (0.0 + v1))}
                _ = v16
            }
        case 32767:
            v17 := []int{42, 898}
            _ = v17
        case 1:
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 float32 = (v1 / (((v1 + v1) - v1) * 1e10))
                v20 := GenS3{F0: 100}
                v21 := []int{(fi18 + (^(fi18)))}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        default:
            switch 2147483647 {
            case 37:
                v1 -= (v1 / v1)
            case 2147483647:
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 int = -609
                    var v24 bool = true
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
        }
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint32 = 0
        v27 := []int{fi25}
        v1 = v1
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 string = "a"
    for fi29 := 0; fi29 < 1; fi29++ {
        fmt.Println(v28, v28)
        var v30 bool = true
        v31 := GenS4{F0: 94}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 uint8 = 0
    var v33 int16 = v3
    if false {
        var v34 uint = 65535
        v35 := GenS1{F0: 0}
        if false {
            var v36 int32 = -606
            v37 := GenS4{F0: 942}
            _ = v36
            _ = v37
        }
        _ = v34
        _ = v35
    }
    fmt.Println(v1)
    var v38 float64 = (float64((v1 - ((v1 + v1) * v1))) - -2.25)
    v39 := []int{32768, 264, 336, 0}
    switch (v28 + v28) {
    case "\n\t":
        var v40 float64 = float64(v1)
        _ = v40
    case "hello":
        fmt.Println(v32)
    case "a":
        var v41 uint64 = 1
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        v2[359] = ""
        if true {
            for fi43 := 0; fi43 < 3; fi43++ {
                var v44 string = "hello"
                var v45 uint32 = 977
                fmt.Println(v33, v44, v1)
                _ = fi43
                _ = v44
                _ = v45
            }
            var v46 int = (fi42 >> ((723) & 63))
            var v47 float32 = 0.0
            _ = v46
            _ = v47
        } else {
            v48 := GenS3{F0: -1}
            fmt.Println(v38, v28, v1)
            v49 := GenS2{F0: 2, F1: v28, F2: ("hello" <= ("a" + v28))}
            _ = v48
            _ = v49
        }
        v50 := GenS4{F0: 942}
        _ = fi42
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v39
}
