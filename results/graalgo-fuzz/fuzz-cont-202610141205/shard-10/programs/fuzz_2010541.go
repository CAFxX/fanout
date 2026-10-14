package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float32 = (0.0 / 1.5)
        _ = fi7
        _ = v8
    }
    var v9 int8 = -112
    var v10 string = ""
    if (0 == 45) {
        switch ("hello" + "") {
        case "a":
            var v11 uint64 = 757
            _ = v11
        }
        v10 = "foo bar"
        fmt.Println(v10)
    }
    v12 := []int{32767, 32768, int(v9)}
    _ = v10
    _ = v12
    _ = v9
    return 1
}

func main() {
    var v1 int8 = -125
    v2 := make(map[string]int)
    var v3 int8 = 95
    v4 := []uint32{}
    var v5 int64 = 415
    var v6 uint32 = 987
    var v13 uint64 = 154
    v1 -= -1
    v14 := v2["\n\t"]
    for fi15 := 0; fi15 < 19; fi15++ {
        v3 = v1
        v14 = (-(fi15))
        _ = fi15
    }
    v3 += 40
    v2["\n\t"] = v14
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 float64 = 0.0
        switch int(v1) {
        case 243:
            v3 = 56
        }
        v1 = v3
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
