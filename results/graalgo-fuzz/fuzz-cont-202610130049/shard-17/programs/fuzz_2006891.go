package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
type GenS3 struct { F0 int8; F1 uint; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint8; F2 int }
type GenS5 struct { F0 uint16; F1 uint; F2 int }
type GenS6 struct { F0 bool }

func main() {
    var v1 int64 = 0
    var v2 uint = 42
    var v3 float32 = float32(1e10)
    var v4 bool = false
    if true {
        if v4 {
            v5 := GenS1{F0: 0}
            v2 = 186
            _ = v5
        }
        if false {
            if (1 <= 4294967295) {
                for fi6 := 0; fi6 < 14; fi6++ {
                    v4 = v4
                    v3 -= 1e10
                    switch (fi6 << ((v2) & 63)) {
                    case 128:
                        v4 = v4
                    }
                    _ = fi6
                }
                v7 := GenS1{F0: 54}
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    v8 := []int{-1, 536, 32768, 4294967295}
                    _ = v8
                case "hello":
                    var v9 uint64 = 18446744073709551615
                    _ = v9
                }
                _ = v7
            }
            if v4 {
                for fi10 := 0; fi10 < 6; fi10++ {
                    v4 = (1 >= 894)
                    v3 = (v3 - v3)
                    _ = fi10
                }
                if true {
                    var v11 uint32 = 137
                    v11 -= v11
                    _ = v11
                } else {
                    v2 = (v2 / ((0) | 1))
                    var v12 uint8 = 1
                    _ = v12
                }
            } else {
                var v13 uint16 = 854
                if ((v1 | -180) != v1) {
                    v1 = (^(v1))
                } else {
                    v14 := GenS1{F0: 0}
                    _ = v14
                }
                _ = v13
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                fmt.Println(v2, v3)
                if v4 {
                    fmt.Println(v3)
                }
                _ = fi15
            }
        }
    }
    var v16 bool = true
    fmt.Println(v3)
    if true {
        v2 = v2
        switch 0 {
        case 1:
            v17 := GenS2{F0: 255, F1: -2147483648, F2: 396}
            _ = v17
        }
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS3{F0: 0, F1: v2, F2: 0}
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v2, v2)
    var v20 uint8 = 0
    v21 := GenS4{F0: v4, F1: (^(v20)), F2: 32767}
    v22 := []int{}
    if false {
        var v23 uint32 = 2
        _ = v23
    }
    v20 = 1
    if true {
        fmt.Println(v20, v16, v3)
    } else {
        if true {
            switch "\n\t" {
            case "a":
                v24 := GenS5{F0: 65535, F1: (v2 & v2), F2: 127}
                _ = v24
            case "":
                if v16 {
                    var v25 uint64 = 1
                    _ = v25
                } else {
                    var v26 int16 = -1
                    _ = v26
                }
            case "hello":
                switch ("" + ((("foo bar" + ("foo bar" + "foo bar")) + "a") + "\n\t")) {
                case "hello":
                    v27 := GenS6{F0: v16}
                    _ = v27
                }
            }
            var v28 uint64 = 18446744073709551615
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
