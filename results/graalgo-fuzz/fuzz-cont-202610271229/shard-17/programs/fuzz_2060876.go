package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 2
    var v3 uint8 = 0
    v2 += (v2 + v2)
    v4 := GenS1{F0: 0}
    var v5 uint32 = 804
    switch ("" + (((("hello" + "") + "") + "") + (("\n\t" + "\n\t") + ""))) {
    case "foo bar":
        delete(v1, (v5 * v5))
    case "a":
        fmt.Println(v3, v3, v2)
    case "hello":
        if (!(!(2 >= 65535))) {
            var v6 int16 = 0
            v6 -= v6
            if true {
                var v7 uint16 = 2
                _ = v7
            } else {
                for fi8 := 0; fi8 < 4; fi8++ {
                    var v9 int16 = (v6 + (v6 << ((v2) & 63)))
                    _ = fi8
                    _ = v9
                }
                v10 := GenS1{F0: int64(v2)}
                fmt.Println(v6)
                _ = v10
            }
            _ = v6
        }
    }
    if (599 >= v2) {
        v3 += v3
        for fi11 := 0; fi11 < 7; fi11++ {
            v5 += v5
            _ = fi11
        }
        v3 = uint8(v5)
    }
    var v12 uint = 0
    var v13 int8 = 40
    v2 -= v2
    switch ("a" + "\n\t") {
    case "":
        if (!false) {
            var v14 int64 = 199
            _ = v14
        }
    }
    var v15 uint8 = 255
    if false {
        v13 -= v13
        var v16 int32 = -2147483648
        _ = v16
    } else {
        if ("hello" != ((("\n\t" + ("a" + "")) + "a") + "hello")) {
            for fi17 := 0; fi17 < 15; fi17++ {
                v13 = v13
                switch "\n\t" {
                case "a":
                    fmt.Println(v13, fi17)
                }
                _ = fi17
            }
        } else {
            if (75 > -1) {
                for fi18 := 0; fi18 < 6; fi18++ {
                    v19 := []int{(-(fi18)), (fi18 % ((fi18) | 1)), 2147483647}
                    _ = fi18
                    _ = v19
                }
            }
        }
    }
    if (-883 == 32767) {
        delete(v1, (^(v5)))
        if (0 >= 1) {
            var v20 int8 = v13
            var v21 uint32 = v5
            _ = v20
            _ = v21
        } else {
            var v22 string = ""
            _ = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
