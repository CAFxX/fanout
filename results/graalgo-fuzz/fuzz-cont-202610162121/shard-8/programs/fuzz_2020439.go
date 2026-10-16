package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("foo bar" + ("foo bar" + "foo bar"))
    var v3 uint32 = 42
    v4 := make(map[int]int)
    v3 += (^(v3))
    v2 = "hello"
    var v5 uint16 = 32767
    v2 = v2
    if true {
        var v6 uint32 = v3
        fmt.Println(v3)
        v3 = v3
        _ = v6
    } else {
        switch (v2 + ("foo bar" + (v2 + v2))) {
        case "\n\t":
            v7 := []int{}
            _ = v7
        case "a":
            v3 = 676
        case "hello":
            var v8 int8 = 0
            _ = v8
        }
    }
    var v9 float64 = 0.1
    if true {
        var v10 float32 = 3.14159
        _ = v10
    }
    if (v9 < v9) {
        switch v2 {
        case "a":
            if true {
                v9 = v9
            }
        case "":
            switch v2 {
            case "\n\t":
                var v11 int = 127
                _ = v11
            case "hello":
                var v12 uint = 255
                _ = v12
            case "":
                v2 = (v2 + ("\n\t" + v2))
            }
        }
        if false {
            var v13 uint64 = 0
            switch ((v2 + v2) + ("" + v2)) {
            case "a":
                fmt.Println(v13)
            case "\n\t":
                v13 = (v13 - v13)
            case "foo bar":
                v1[-80] = (!(uint(v5) > 0))
            }
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 int = (fi14 - 9223372036854775807)
                if (0 != -1) {
                    var v16 int8 = -128
                    _ = v16
                } else {
                    fmt.Println(fi14)
                    v3 = 133
                    var v17 int64 = 0
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
            _ = v13
        }
    } else {
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS1{F0: (415 < 4294967295), F1: -58, F2: 738}
            var v20 int64 = 1
            if (413 < 379) {
                v21 := GenS2{F0: v2}
                _ = v21
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        var v22 int16 = 32
        v23 := GenS1{F0: true, F1: -1, F2: 42}
        _ = v22
        _ = v23
    }
    var v24 uint8 = 127
    var v25 float32 = (-2.25 + 0.0)
    if false {
        if true {
            if (14 > -1) {
                v2 = "\n\t"
                var v26 uint8 = 2
                _ = v26
            } else {
                if true {
                    v27 := GenS2{F0: ((v2 + (v2 + (v2 + ("\n\t" + (v2 + "hello"))))) + (v2 + ((v2 + v2) + ("\n\t" + (("" + (v2 + v2)) + v2)))))}
                    v28 := GenS3{F0: int64(v24), F1: v24}
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    v30 := GenS1{F0: (42 > 42), F1: int8(fi29), F2: -9223372036854775808}
                    _ = fi29
                    _ = v30
                }
            }
            if (((true || (!false)) || ((!true) && ((false || (v2 > ("a" + v2))) && false))) && false) {
                v31 := []int{4294967295, -1}
                _ = v31
            }
            v24 = 42
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
