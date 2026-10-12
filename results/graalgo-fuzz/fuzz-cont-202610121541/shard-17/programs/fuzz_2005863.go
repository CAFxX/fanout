package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint8; F2 int32 }
func genFn0(p0 string) int16 {
    if (!(false || ((!false) || false))) {
        var v4 bool = false
        _ = v4
    }
    if false {
        v5 := GenS1{F0: 32767, F1: true}
        _ = v5
    }
    _ = p0
    return 449
}

func main() {
    var v1 bool = true
    var v2 uint8 = 127
    v3 := make(map[string]int)
    delete(v3, ("a" + "hello"))
    v2 = 1
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{}
        delete(v3, (("hello" + ("foo bar" + "hello")) + ((("\n\t" + ("a" + "a")) + "a") + "foo bar")))
        _ = fi6
        _ = v7
    }
    if (32767 <= -1) {
        if v1 {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{(fi8 / ((fi8) | 1))}
                v10 := GenS2{F0: 920, F1: fi8}
                _ = fi8
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 13; fi11++ {
                for fi12 := 0; fi12 < 14; fi12++ {
                    v13 := []int{32768, (fi11 + fi11)}
                    var v14 int8 = 0
                    var v15 int8 = (-(v14))
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                var v16 int8 = 41
                delete(v3, ("" + ((("" + (("a" + (("" + "\n\t") + ("" + "hello"))) + "hello")) + (("foo bar" + ((("\n\t" + "") + ("a" + "hello")) + ("foo bar" + "a"))) + "\n\t")) + "foo bar")))
                _ = fi11
                _ = v16
            }
            fmt.Println(v2, v2)
        } else {
            v17 := GenS2{F0: -421, F1: 127}
            switch ("hello" + "\n\t") {
            case "foo bar":
                var v18 uint16 = 931
                _ = v18
            }
            if ((v2 ^ (v2 >> ((4294967295) & 63))) > (v2 | (v2 * v2))) {
                fmt.Println(v1, v2)
                v19 := GenS3{F0: 2, F1: -9223372036854775808, F2: 168}
                _ = v19
            } else {
                v1 = (!(v1 && (!(genFn0("\n\t") == 42))))
                v20 := GenS4{F0: (2147483647 >= 0), F1: v2, F2: -349}
                _ = v20
            }
            _ = v17
        }
        var v21 int8 = 127
        _ = v21
    }
    fmt.Println(v1, v1)
    v22 := []int{-1, 0}
    v22 = append(v22, -2147483648)
    var v23 string = "hello"
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    v2 += (v2 + v2)
    if len(v22) > 0 {
        v22[0] = 9223372036854775807
    }
    v23 = (((v23 + (("" + v23) + "foo bar")) + ("a" + ("\n\t" + ""))) + ("hello" + v23))
    var v24 int8 = 100
    v2 -= v2
    v2 += 128
    var v25 uint64 = 1
    v25 += (v25 >> ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
