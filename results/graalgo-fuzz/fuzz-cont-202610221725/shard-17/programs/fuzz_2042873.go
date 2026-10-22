package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint16; F2 int64 }
func genFn0(p0 int16, p1 bool, p2 bool) int64 {
    if (!(1 != -1)) {
        p1 = (!(0 == -1))
        var v5 float64 = -2.25
        _ = v5
    }
    p2 = (971 != 121)
    var v6 uint64 = 4294967295
    p0 = (-(p0))
    v7 := []int{0, 1, 659}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -940
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    if (!(3.14159 > 1.5)) {
        var v8 bool = false
        v8 = (!v8)
        var v9 bool = ((0.0 - (0.1 - -2.25)) == (((float64(3.14159) - 1e10) / ((float64(1.5) / 1e10) * ((((3.14159 - 1.5) + 1e10) + 1.5) + ((-2.25 - (0.0 + (1.5 / 1.5))) - ((0.0 / 1e10) - -2.25))))) + 0.0))
        _ = v8
        _ = v9
    } else {
        switch ("hello" + "\n\t") {
        case "a":
            var v10 bool = false
            _ = v10
        }
    }
    v1 = (v1 << ((v1) & 63))
    if false {
        v1 -= 65535
        fmt.Println(v1)
    }
    v11 := GenS1{F0: 119, F1: -9223372036854775808}
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS2{F0: 65535, F1: 0}
        v3[2] = (fi12 / ((fi12) | 1))
        _ = fi12
        _ = v13
    }
    if (-9223372036854775808 >= 1) {
        v1 = 2
        if true {
            fmt.Println(v1)
            v14 := []int{128, -9223372036854775808, 32767, 4294967295}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v14
        }
        v1 += (^(v1))
    } else {
        var v15 uint = v1
        v16 := []int{}
        _ = v15
        _ = v16
    }
    var v17 uint64 = 4294967295
    fmt.Println(v1, v17)
    if (3.14159 >= 0.0) {
        fmt.Println(v17)
        v18 := GenS2{F0: 0, F1: v17}
        if (("foo bar" == ("hello" + "hello")) || false) {
            fmt.Println(v17)
            v19 := GenS3{F0: (v1 & 0), F1: 255, F2: -768}
            switch 32768 {
            case 2147483647:
                v1 = v1
            }
            _ = v19
        }
        _ = v18
    } else {
        _ = len(v4)
        delete(v4, 2147483647)
    }
    delete(v3, 42)
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 float64 = 1.5
        _ = fi20
        _ = v21
    }
    var v22 float32 = 0.1
    if (98 == 128) {
        if true {
            v23 := []int{-2147483648, -538, 191}
            var v24 uint32 = 4294967295
            var v25 uint8 = 2
            _ = v23
            _ = v24
            _ = v25
        }
    }
    v26 := GenS2{F0: 32767, F1: v17}
    v22 = 0.0
    fmt.Println(v17)
    v17 += v17
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
}
