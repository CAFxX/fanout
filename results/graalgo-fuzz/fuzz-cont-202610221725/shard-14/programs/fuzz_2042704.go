package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() bool {
    v4 := GenS1{F0: 65535}
    var v5 float32 = 3.14159
    v6 := []int{4294967295, -9223372036854775808, 42, 770}
    _ = v4
    _ = v5
    _ = v6
    return (4294967295 >= 255)
}

func main() {
    var v1 string = "a"
    var v2 string = v1
    v3 := []uint32{895, 4294967295, 4294967295}
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = "foo bar"
        var v7 uint8 = 128
        _ = v7
    }
    v1 = v1
    var v8 uint64 = 2
    v8 -= (^(v8))
    fmt.Println(v1, v1, v1)
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint = 0
        _ = fi9
        _ = v10
    }
    if len(v3) > 0 {
        v3[0] = 65535
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v11 := GenS1{F0: 322}
    v1 = "hello"
    switch 32768 {
    case 9223372036854775807:
        var v12 string = v1
        _ = v12
    case 128:
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 13; fi14++ {
                v3 = append(v3, 2)
                fmt.Println(v2, v8, v2)
                fmt.Println(fi14, fi14)
                v2 = "foo bar"
                if len(v3) > 0 {
                    v3[0] = 0
                }
                _ = fi14
            }
            switch (v1 + v1) {
            case "hello":
                var v15 float32 = 0.0
                _ = v15
            case "\n\t":
                for fi16 := 0; fi16 < 12; fi16++ {
                    v2 = (v1 + "")
                    var v17 int16 = -32768
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi13
        }
    }
    var v18 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
