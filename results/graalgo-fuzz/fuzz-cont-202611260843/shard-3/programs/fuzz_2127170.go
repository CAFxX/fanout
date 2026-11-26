package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint }
type GenS4 struct { F0 uint; F1 uint; F2 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 int; F2 uint64 }
type GenS7 struct { F0 uint }

func main() {
    var v1 string = ("a" + "a")
    var v2 int = 128
    var v3 int8 = 42
    var v4 int8 = v3
    if true {
        v5 := GenS1{F0: 1}
        if ((0 == 2) || (!(2 > 42))) {
            var v6 uint32 = 2147483648
            switch v1 {
            case "a":
                v6 -= 2147483648
            case "foo bar":
                fmt.Println(v1)
            default:
                var v7 float64 = 1e10
                _ = v7
            }
            fmt.Println(v1, v3)
            _ = v6
        } else {
            v2 += 127
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                v3 = (-(v4))
                v10 := GenS2{F0: 2147483647, F1: (((v1 + (v1 + v1)) + v1) + ("" + (("hello" + v1) + "foo bar"))), F2: 214}
                var v11 bool = (-548 >= 0)
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        _ = v5
    } else {
        v3 += int8(v2)
        fmt.Println(v1, v3, v2)
        var v12 uint = 465
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := []int{fi13, fi13, v2}
        v2 = v2
        _ = fi13
        _ = v14
    }
    var v15 float32 = (-2.25 - float32((1e10 / ((float64(0.0) * -2.25) + 1e10))))
    v16 := GenS1{F0: 703}
    var v17 int64 = 1
    v18 := GenS3{F0: 595, F1: (v3 | (v4 << ((18446744073709551615) & 63))), F2: 0}
    var v19 int8 = (v4 << ((4294967295) & 63))
    v20 := GenS4{F0: 18446744073709551615, F1: 4294967295, F2: 65535}
    if false {
        v4 = (v3 - v4)
        var v21 uint = 1
        var v22 string = "hello"
        _ = v21
        _ = v22
    } else {
        switch v1 {
        case "hello":
            v23 := []int{}
            _ = v23
        default:
            var v24 uint8 = 127
            _ = v24
        }
        switch ((((v1 + "\n\t") + (((v1 + ((v1 + "foo bar") + (v1 + v1))) + "hello") + "\n\t")) + v1) + v1) {
        case "":
            fmt.Println(v3, v3)
        }
    }
    if (127 != 41) {
        var v25 uint64 = 2
        _ = v25
    } else {
        v4 = (v4 % ((100) | 1))
        v19 += (v3 - (v3 >> ((835) & 63)))
    }
    v26 := GenS3{F0: 1, F1: 1, F2: 0}
    fmt.Println(v1, v3, v15)
    for fi27 := 0; fi27 < 1; fi27++ {
        v15 = v15
        for fi28 := 0; fi28 < 10; fi28++ {
            v29 := GenS5{F0: 42}
            if true {
                v17 += (v17 & v17)
                switch -9223372036854775808 {
                case 128:
                    switch -2147483648 {
                    case 42:
                        var v30 uint64 = 18446744073709551615
                        _ = v30
                    case 0:
                        switch (fi28 & fi28) {
                        case 486:
                            var v31 int32 = -2147483648
                            _ = v31
                        default:
                            fmt.Println(v15)
                        }
                    }
                case 0:
                    v32 := GenS6{F0: 845, F1: fi27, F2: 4294967295}
                    _ = v32
                }
            } else {
                v33 := GenS7{F0: 2}
                _ = v33
            }
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    var v34 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v34
    _ = v4
}
