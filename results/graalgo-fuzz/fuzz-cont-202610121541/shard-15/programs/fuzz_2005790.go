package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
func genFn0() uint32 {
    if true {
        var v4 string = ("a" + "\n\t")
        _ = v4
    }
    var v5 string = "a"
    switch 32767 {
    case -9223372036854775808:
        v5 = v5
    case 883:
        fmt.Println(v5)
    case 9223372036854775807:
        var v6 float32 = 1e10
        _ = v6
    }
    v5 = ""
    _ = v5
    return 65535
}
func genFn1(p0 int) int8 {
    var v7 int32 = -2147483648
    fmt.Println(p0, p0)
    for fi8 := 0; fi8 < 3; fi8++ {
        if (127 != 1) {
            v7 = (v7 | v7)
        }
        _ = fi8
    }
    v9 := GenS1{F0: 0}
    _ = p0
    _ = v7
    _ = v9
    return 127
}

func main() {
    var v1 int64 = 411
    var v2 uint32 = 107
    var v3 uint32 = (^(v2))
    for fi10 := 0; fi10 < 20; fi10++ {
        if true {
            v1 = -942
        } else {
            v11 := GenS1{F0: 255}
            v1 = 9223372036854775807
            _ = v11
        }
        v12 := []int{}
        var v13 int64 = -9223372036854775808
        _ = fi10
        _ = v12
        _ = v13
    }
    fmt.Println(v1)
    var v14 string = "\n\t"
    if true {
        v2 = v2
        var v15 int = 1
        fmt.Println(v15)
        _ = v15
    }
    v16 := GenS2{F0: (v1 >> ((1) & 63))}
    if (((v2 + v2) <= (v2 | v3)) || (((127 >= 6) || (true && (0 >= 814))) && false)) {
        v17 := GenS2{F0: v1}
        v14 = "foo bar"
        v18 := GenS2{F0: 0}
        _ = v17
        _ = v18
    }
    var v19 int = -1
    v1 = v1
    var v20 int16 = 32767
    switch "" {
    case "a":
        v21 := []int{v19}
        _ = v21
    case "hello":
        if false {
            v1 = (v1 * 1)
            var v22 bool = (!false)
            for fi23 := 0; fi23 < 1; fi23++ {
                v24 := GenS3{F0: 127}
                _ = fi23
                _ = v24
            }
            _ = v22
        }
    case "":
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS3{F0: 69}
            _ = fi25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        v28 := GenS3{F0: genFn1((fi27 >> ((255) & 63)))}
        _ = fi27
        _ = v28
    }
    if true {
        v29 := GenS1{F0: 223}
        _ = v29
    }
    switch v19 {
    case 32768:
        for fi30 := 0; fi30 < 9; fi30++ {
            v2 = (^(v3))
            _ = fi30
        }
    case 9223372036854775807:
        v14 = "a"
    case -528:
        fmt.Println(v2, v20)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
