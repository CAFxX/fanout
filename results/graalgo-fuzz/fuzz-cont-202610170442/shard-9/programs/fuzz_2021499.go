package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
func genFn0(p0 uint) uint {
    var v6 uint32 = 2
    v7 := GenS1{F0: 1, F1: 9223372036854775807, F2: -643}
    if (18446744073709551615 <= 4294967295) {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int16 = 32767
            p0 = p0
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return p0
}
func genFn1() bool {
    v10 := []int{241, 2147483647, 1, 4294967295}
    if len(v10) > 0 {
        v10[0] = 4294967295
    }
    _ = v10
    return (127 != -9223372036854775808)
}

func main() {
    var v1 uint64 = 1
    var v2 int64 = 9223372036854775807
    var v3 uint8 = 1
    v4 := []uint32{42, 42, 65535}
    v5 := []string{""}
    fmt.Println(v1, v1)
    v3 = 127
    if len(v5) > 0 {
        v5[0] = ("foo bar" + ("foo bar" + ""))
    }
    v1 = v1
    if false {
        var v11 string = "foo bar"
        for fi12 := 0; fi12 < 9; fi12++ {
            v2 = (v2 * (v2 % (((^(v2))) | 1)))
            if ((v11 > "hello") || ((^(v3)) == 127)) {
                v11 = (v11 + v11)
                if len(v5) > 0 {
                    v5[0] = ("a" + "\n\t")
                }
            } else {
                var v13 uint16 = 42
                v4 = append(v4, 230)
                _ = v13
            }
            _ = fi12
        }
        var v14 uint32 = 0
        _ = v11
        _ = v14
    }
    fmt.Println(v2)
    var v15 uint = 2
    if true {
        v5 = append(v5, ((((("foo bar" + ("a" + ("" + (("a" + ((("\n\t" + ("a" + (("\n\t" + ("" + "")) + ""))) + (("" + (("a" + ((("" + (("\n\t" + "") + "\n\t")) + ("foo bar" + ("foo bar" + ("hello" + "hello")))) + "a")) + "")) + "foo bar")) + ("" + "a"))) + ("hello" + (("\n\t" + "\n\t") + "\n\t")))))) + ("" + "a")) + "hello") + "\n\t") + "hello"))
    } else {
        v2 += v2
    }
    var v16 int16 = 1000
    if (32768 > -1) {
        var v17 int16 = 397
        _ = v17
    }
    v18 := GenS1{F0: v2, F1: 4294967295, F2: 0}
    var v19 string = ("foo bar" + "\n\t")
    v3 = uint8(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
