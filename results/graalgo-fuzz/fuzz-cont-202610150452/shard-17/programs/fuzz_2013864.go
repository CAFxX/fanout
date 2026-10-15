package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int8; F1 bool; F2 int8 }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]bool)
    var v3 uint32 = 42
    v4 := []string{("hello" + "a"), ""}
    v1 += (^(v1))
    var v5 uint32 = v3
    v5 += v3
    switch "\n\t" {
    case "":
        var v6 int16 = 0
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 2; fi7++ {
            switch (^(fi7)) {
            case 134:
                var v8 uint = uint(v1)
                _ = v8
            case -1:
                var v9 uint64 = 0
                _ = v9
            }
            switch int(v3) {
            case 42:
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 uint64 = 4294967295
                    fmt.Println(fi10)
                    v12 := GenS1{F0: -164}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case 0:
                v3 = v3
            default:
                if false {
                    fmt.Println(v5, fi7)
                    switch 806 {
                    case 128:
                        if true {
                            if ((128 >= v1) && true) {
                                var v13 int8 = 11
                                for fi14 := 0; fi14 < 19; fi14++ {
                                    v3 = (v3 & (v5 & v3))
                                    v13 += -1
                                    _ = fi14
                                }
                                for fi15 := 0; fi15 < 8; fi15++ {
                                    fmt.Println(v3, fi15)
                                    _ = fi15
                                }
                                _ = v13
                            }
                        } else {
                            for fi16 := 0; fi16 < 6; fi16++ {
                                v1 = 255
                                _ = fi16
                            }
                            v4 = append(v4, "hello")
                        }
                    }
                }
            }
            v17 := GenS1{F0: 1}
            _ = fi7
            _ = v17
        }
    }
    var v18 bool = (65535 <= 2)
    if (!v18) {
        v19 := GenS2{F0: -873, F1: v3}
        if v18 {
            v20 := v2[4294967295]
            _ = v20
        } else {
            switch -1 {
            case 0:
                var v21 int8 = 42
                _ = v21
            }
        }
        v18 = v18
        _ = v19
    } else {
        v1 -= 2
    }
    switch "\n\t" {
    case "\n\t":
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 uint16 = 42
            _ = fi22
            _ = v23
        }
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        v24 := GenS3{F0: 1, F1: v18, F2: 1}
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 int8 = 127
        _ = len(v4)
        _ = fi25
        _ = v26
    }
    v27 := GenS4{F0: "hello", F1: -2147483648}
    switch -135 {
    case -2147483648:
        var v28 int8 = -128
        _ = v28
    default:
        for fi29 := 0; fi29 < 16; fi29++ {
            v18 = ((v1 + (v1 << ((0) & 63))) <= 190)
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 uint32 = v5
                var v32 int = (fi30 & (fi29 | fi30))
                _ = len(v4)
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi29
        }
    }
    var v33 string = "\n\t"
    v1 += 2
    var v34 bool = v18
    var v35 uint8 = v1
    var v36 int = 0
    fmt.Println(v34, v34, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
