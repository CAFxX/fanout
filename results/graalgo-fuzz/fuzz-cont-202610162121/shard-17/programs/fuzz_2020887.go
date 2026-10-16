package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
func genFn0(p0 bool, p1 uint8, p2 uint32) uint8 {
    fmt.Println(p2)
    if ((42 >= -918) || (1e10 == 3.14159)) {
        v7 := []int{-9223372036854775808, 0, 127}
        p0 = p0
        _ = v7
    } else {
        if (p2 == (^(p2))) {
            fmt.Println(p1)
            p1 += p1
        }
    }
    if (-2.25 >= 0.1) {
        var v8 int = -797
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 - p1)
}

func main() {
    var v1 int16 = 65
    var v2 uint8 = 1
    var v3 int16 = v1
    var v4 int8 = -40
    var v5 uint = 698
    v6 := make(map[uint32]bool)
    v6[65535] = false
    var v9 int64 = 0
    var v10 uint64 = 42
    v1 = 505
    fmt.Println(v5)
    var v11 uint16 = 205
    fmt.Println(v2, v10)
    if (42 >= 2147483647) {
        v12 := GenS1{F0: 42, F1: true}
        fmt.Println(v5, v10)
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 bool = true
                _ = fi14
                _ = v15
            }
            var v16 uint16 = v11
            _ = fi13
            _ = v16
        }
        _ = v12
    }
    v17 := GenS1{F0: 1, F1: false}
    if (genFn0((true || (true && true)), (v2 >> ((v5) & 63)), 352) < v2) {
        if true {
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 int16 = int16(v11)
                _ = fi18
                _ = v19
            }
            var v20 uint16 = 255
            _ = v20
        } else {
            v21 := GenS1{F0: (v4 ^ (v4 * v4)), F1: false}
            _ = v21
        }
        v22 := []int{}
        _ = v22
    } else {
        v23 := GenS1{F0: -128, F1: true}
        _ = v23
    }
    switch 177 {
    case 42:
        var v24 int = -842
        _ = v24
    case 0:
        switch int(v9) {
        case 4294967295:
            v1 = 674
        }
    case -35:
        for fi25 := 0; fi25 < 20; fi25++ {
            if false {
                v26 := []int{2147483647}
                _ = v26
            }
            var v27 string = "\n\t"
            _ = fi25
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        for fi29 := 0; fi29 < 18; fi29++ {
            switch ("hello" + "a") {
            case "a":
                fmt.Println(v11, fi28)
                var v30 uint16 = 65535
                fmt.Println(v2, v11, v1)
                _ = v30
            case "\n\t":
                if false {
                    var v31 string = "hello"
                    _ = v31
                }
            }
            var v32 uint64 = (v10 ^ (v10 >> (((v5 ^ v5)) & 63)))
            _ = fi29
            _ = v32
        }
        if (!true) {
            for fi33 := 0; fi33 < 17; fi33++ {
                fmt.Println(v10, v9)
                v5 += 4294967295
                switch fi28 {
                case 2147483647:
                    v10 = v10
                }
                _ = fi33
            }
        } else {
            v4 -= -128
        }
        fmt.Println(v4, v1, v10)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
