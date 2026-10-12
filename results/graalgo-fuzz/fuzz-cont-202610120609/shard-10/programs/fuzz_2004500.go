package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 string; F1 uint64 }

func main() {
    var v1 uint = 42
    var v2 uint64 = 286
    var v3 uint = v1
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 844
    var v6 int8 = -1
    v1 -= 2
    v1 = v1
    var v7 bool = (-2147483648 > 42)
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: int(v3), F1: v6}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: -1, F1: (v6 ^ v6)}
    var v11 string = (("\n\t" + ("foo bar" + ("" + "hello"))) + "hello")
    var v12 float64 = 0.1
    v7 = (1 > 1000)
    switch (v11 + "a") {
    case "hello":
        var v13 bool = ((!v7) || v7)
        _ = v13
    case "\n\t":
        v6 = v6
    default:
        for fi14 := 0; fi14 < 1; fi14++ {
            switch "\n\t" {
            case "":
                if (v6 >= (v6 + (v6 * v6))) {
                    switch (((v11 + v11) + v11) + ("" + ((v11 + v11) + ""))) {
                    case "\n\t":
                        v12 += v12
                    case "":
                        if v7 {
                            v12 += (v12 - (0.0 + v12))
                            var v15 float32 = (1.5 / (((-2.25 - (-2.25 - float32(v12))) + 1.5) / float32(v12)))
                            for fi16 := 0; fi16 < 19; fi16++ {
                                v7 = v7
                                _ = fi16
                            }
                            _ = v15
                        } else {
                            v17 := []int{fi14, 1}
                            _ = v17
                        }
                    case "a":
                        for fi18 := 0; fi18 < 15; fi18++ {
                            v19 := GenS1{F0: int(v1), F1: v6}
                            _ = fi18
                            _ = v19
                        }
                    default:
                        v20 := GenS2{F0: 837, F1: -1}
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 19; fi21++ {
                        v22 := GenS3{F0: "hello", F1: 42}
                        _ = fi21
                        _ = v22
                    }
                    var v23 int64 = int64(v2)
                    _ = v23
                }
            case "hello":
                v24 := GenS1{F0: -374, F1: v6}
                _ = v24
            }
            _ = fi14
        }
    }
    v25 := []int{-567, -776}
    if false {
        var v26 string = "foo bar"
        _ = v26
    }
    var v27 int32 = 1
    var v28 uint8 = 127
    if v7 {
        for fi29 := 0; fi29 < 1; fi29++ {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            var v30 string = v11
            if true {
                var v31 uint8 = v28
                _ = v31
            }
            _ = fi29
            _ = v30
        }
    } else {
        switch 2147483647 {
        case -768:
            for fi32 := 0; fi32 < 10; fi32++ {
                var v33 bool = v7
                for fi34 := 0; fi34 < 16; fi34++ {
                    var v35 uint32 = (^(v5))
                    v36 := GenS2{F0: int64(v1), F1: 1000}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                var v37 int = fi32
                _ = fi32
                _ = v33
                _ = v37
            }
        case 0:
            for fi38 := 0; fi38 < 10; fi38++ {
                v3 = v3
                if (!((!(2 != uint16(v1))) || v7)) {
                    var v39 bool = true
                    var v40 uint64 = v2
                    _ = v39
                    _ = v40
                } else {
                    v5 -= 1
                }
                _ = fi38
            }
        case 127:
            v7 = (!v7)
        }
    }
    if (v12 > 1e10) {
        if len(v25) > 0 {
            v25[0] = 4294967295
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
