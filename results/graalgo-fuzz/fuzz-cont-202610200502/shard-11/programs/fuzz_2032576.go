package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() string {
    v7 := []int{}
    v8 := GenS1{F0: 72}
    var v9 float32 = ((1e10 * 3.14159) + ((0.1 + (1.5 - float32(1.5))) - (3.14159 - ((1.5 * (0.0 + 0.0)) * (1e10 + (float32(float64((3.14159 / 1.5))) + 0.1))))))
    _ = v7
    _ = v8
    _ = v9
    return ("\n\t" + "foo bar")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 int = 4294967295
    v4 := make(map[uint32]int)
    var v5 int8 = 0
    var v6 string = ""
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 bool = (0 < -705)
        v12 := GenS1{F0: 708}
        fmt.Println(v11, v5, v6)
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        fmt.Println(v6)
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint64 = 0
            v16 := []int{fi13, 0, (fi14 | v3)}
            v17 := GenS1{F0: 65535}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    var v18 bool = false
    var v19 uint64 = 2
    delete(v2, 637)
    fmt.Println(v18, v3)
    var v20 uint64 = v19
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 float64 = float64(-2.25)
        v3 -= 128
        v20 -= v19
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 42}
    fmt.Println(v3, v20)
    for fi24 := 0; fi24 < 9; fi24++ {
        if (v18 && (-1 >= -2147483648)) {
            var v25 uint = 195
            _ = v25
        }
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 uint16 = 207
            _ = fi26
            _ = v27
        }
        var v28 uint8 = 1
        _ = fi24
        _ = v28
    }
    if v18 {
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 uint32 = 1
            _ = fi29
            _ = v30
        }
        var v31 int16 = -1
        _ = v31
    }
    if (true && v18) {
        fmt.Println(v6, v18, v18)
    } else {
        var v32 uint64 = (v19 >> ((18446744073709551615) & 63))
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
