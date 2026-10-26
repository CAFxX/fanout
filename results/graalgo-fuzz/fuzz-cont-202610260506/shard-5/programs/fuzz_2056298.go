package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() bool {
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: false}
        var v9 float32 = (float32(1e10) / -2.25)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: (!(-32768 >= 42))}
    if true {
        var v11 uint32 = 65535
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 float32 = 3.14159
            var v14 uint8 = 128
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := []int{32767, -1, 693, 42}
        _ = v11
        _ = v15
    }
    var v16 string = ("a" + ((("\n\t" + "foo bar") + "a") + ""))
    _ = v10
    _ = v16
    return (!((42 <= 2147483648) || (-2.25 >= (0.0 / ((-2.25 * (1.5 * (((((float32((0.1 * 3.14159)) / (1e10 / (3.14159 - 1.5))) - (((3.14159 + (float32(float64(1.5)) * -2.25)) + (0.0 + 1.5)) - float32(((0.0 - 0.0) / (1e10 / (1e10 + ((0.0 / (1e10 / 3.14159)) + -2.25))))))) * 3.14159) + ((1e10 - -2.25) * 0.1)) / 1.0))) / 3.14159)))))
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]int)
    var v3 uint64 = 2
    var v4 int8 = -128
    var v5 uint64 = (v3 / ((42) | 1))
    var v6 uint16 = 235
    switch ("\n\t" + ("foo bar" + ("" + "a"))) {
    case "a":
        v17 := GenS1{F0: true}
        _ = v17
    default:
        v1 += (v1 + -9223372036854775808)
    }
    v2[int(v1)] = 0
    var v18 int16 = 1000
    var v19 int32 = -505
    v20 := GenS1{F0: true}
    if true {
        v6 = v6
    } else {
        var v21 uint = 4294967295
        if (true && (0.0 <= -2.25)) {
            var v22 float32 = float32((1e10 + 0.1))
            for fi23 := 0; fi23 < 18; fi23++ {
                v22 = v22
                _ = fi23
            }
            var v24 int32 = 751
            _ = v22
            _ = v24
        }
        _ = v21
    }
    var v25 string = ""
    v26 := []int{9223372036854775807, -43, 32768}
    _ = len(v26)
    for fi27 := 0; fi27 < 15; fi27++ {
        switch v25 {
        case "foo bar":
            var v28 uint64 = 1
            _ = v28
        case "":
            v29 := v2[(fi27 % ((fi27) | 1))]
            _ = v29
        case "\n\t":
            var v30 uint16 = (v6 >> ((92) & 63))
            _ = v30
        }
        var v31 uint64 = (v3 + 1)
        v32 := GenS1{F0: (v1 >= v1)}
        _ = fi27
        _ = v31
        _ = v32
    }
    var v33 uint = 2
    var v34 int64 = 9223372036854775807
    var v35 int32 = (v19 << ((v33) & 63))
    var v36 uint = v33
    switch v25 {
    case "\n\t":
        var v37 uint64 = (v5 >> ((v36) & 63))
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
