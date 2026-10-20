package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() uint16 {
    var v5 uint8 = 42
    v6 := GenS1{F0: "a"}
    fmt.Println(v5)
    var v7 int64 = -9223372036854775808
    _ = v5
    _ = v6
    _ = v7
    return 266
}

func main() {
    var v1 uint16 = 255
    var v2 float32 = 1e10
    v3 := make(map[uint32]bool)
    var v4 uint64 = 175
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 string = "foo bar"
        fmt.Println(v9)
        _ = fi8
        _ = v9
    }
    var v10 bool = (2147483648 > 4294967295)
    for fi11 := 0; fi11 < 19; fi11++ {
        v2 -= 3.14159
        _ = fi11
    }
    v10 = v10
    v1 -= (v1 % (((v1 & v1)) | 1))
    v2 -= ((float32(float64(0.0)) - ((0.1 + v2) * (0.0 - (v2 + v2)))) * v2)
    if true {
        v4 += v4
    } else {
        fmt.Println(v2)
        var v12 uint16 = 255
        _ = v12
    }
    v4 -= 0
    _ = len(v3)
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 float32 = v2
        if false {
            switch ("foo bar" + "hello") {
            case "hello":
                v15 := GenS1{F0: "\n\t"}
                _ = v15
            }
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                fmt.Println(v10)
                fmt.Println(fi13)
                _ = fi16
            }
            v1 = v1
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            if (-9223372036854775808 != 0) {
                var v18 uint16 = 0
                if (!(127 >= 127)) {
                    v19 := []int{fi13, (fi13 ^ fi17), fi13, (fi13 % ((32767) | 1))}
                    v20 := GenS1{F0: ((("foo bar" + ("" + "")) + "a") + (("foo bar" + ("foo bar" + "\n\t")) + ""))}
                    _ = v19
                    _ = v20
                } else {
                    v3[2] = (!true)
                    v14 = v2
                    var v21 float64 = (0.0 + 3.14159)
                    var v22 uint32 = 225
                    _ = v21
                    _ = v22
                }
                _ = v18
            } else {
                switch "a" {
                case "a":
                    v2 = ((v14 + v2) * float32(float64((float32((3.14159 + 1.5)) / 1e10))))
                }
            }
            _ = fi17
        }
        _ = fi13
        _ = v14
    }
    var v23 uint16 = (^(v1))
    if v10 {
        v24 := []int{}
        var v25 uint16 = (v1 << ((1) & 63))
        _ = v24
        _ = v25
    }
    switch "a" {
    case "a":
        v26 := []int{-290, 912, -612}
        _ = v26
    default:
        for fi27 := 0; fi27 < 17; fi27++ {
            for fi28 := 0; fi28 < 4; fi28++ {
                fmt.Println(v4, v2, v1)
                var v29 int64 = -1
                _ = fi28
                _ = v29
            }
            var v30 int64 = 9223372036854775807
            _ = fi27
            _ = v30
        }
    }
    var v31 int16 = 83
    var v32 int32 = 235
    v4 = v4
    v33 := GenS1{F0: ""}
    v34 := GenS1{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
